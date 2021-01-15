package middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// JWTClaims data structure for claims
type JWTClaims struct {
	AccountID   int64  `json:"account_id"`
	AccountRole string `json:"role_id"`
	PartnerType string `json:"partner_type"`
	Username    string `json:"username"`
	Super       bool   `json:"super"`
	Type        string `json:"type"`
	jwt.StandardClaims
}

// JSONFailed ...
type JSONFailed struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// VerifyExpiresAt function, will override original VerifyExpiresAt function
func (c *JWTClaims) VerifyExpiresAt(cmp int64, req bool) bool {
	var leeway int64 = 60 // one minutes
	return c.StandardClaims.VerifyExpiresAt(cmp-leeway, req)
}

// JWTVerify function to verify json web token
func JWTVerify(rsaPublicKey *rsa.PublicKey, mustAuthorized bool, accountType []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if os.Getenv("NO_TOKEN") == "1" {
				return next(c)
			}

			req := c.Request()
			header := req.Header
			auth := header.Get("Authorization")

			if len(auth) <= 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is empty")
			}

			splitToken := strings.Split(auth, " ")
			if len(splitToken) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is empty")
			}

			if splitToken[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is invalid")
			}

			tokenStr := splitToken[1]
			token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return rsaPublicKey, nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			if claims, ok := token.Claims.(*JWTClaims); token.Valid && ok {
				c.Set("token", token)
				c.Set("tokenStr", tokenStr)
				c.Set("accountID", claims.AccountID)
				c.Set("accountRole", claims.AccountRole)
				c.Set("partnerType", claims.PartnerType)
				c.Set("type", claims.Type)

				if len(accountType) > 0 {
					for _, val := range accountType {
						if claims.Type != val {
							return echo.NewHTTPError(http.StatusUnauthorized, `Unauthorize`)
						}
					}
				}

				return next(c)
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				var errorStr string
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					errorStr = "Token has been expired"
				} else {
					errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
				}
				return echo.NewHTTPError(http.StatusUnauthorized, errorStr)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unknown token error")
			}
		}
	}
}

func VerifyToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			response := new(JSONFailed)

			tokenStr := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
			if len(tokenStr) == 0 {
				response.Success = false
				response.Message = "Not Found Authorization"
				response.Code = http.StatusUnauthorized

				return c.JSON(http.StatusNonAuthoritativeInfo, response)
			}

			claims := jwt.MapClaims{}
			jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			})

			// do something with decoded claims
			if claims["authorised"] == false && claims["email"] == "" {
				response.Success = false
				response.Message = "Unauthorized"
				response.Code = http.StatusUnauthorized

				return c.JSON(http.StatusNonAuthoritativeInfo, response)
			}
			return next(c)
		}
	}
}
