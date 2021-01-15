package middleware

import (
	"github.com/labstack/echo"
	"net/http"
)



func RoleValidation(role []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			accountRole  := c.Get("type")
			if accountRole == `` {
				return echo.NewHTTPError(http.StatusUnauthorized, `Unauthorize`)
			}
			for _, val := range role {
				if accountRole != val {
					return echo.NewHTTPError(http.StatusUnauthorized, `Unauthorize`)
				}
			}
			return next(c)
		}
	}
}