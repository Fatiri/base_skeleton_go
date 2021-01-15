package v1

import (
	"crypto/rsa"

	"github.com/base_skeleton_go/shared"
	"github.com/base_skeleton_go/src/usecase"
	"github.com/labstack/echo"
)

// AccountDelivery ...
type AccountDelivery struct {
	AccountUc usecase.Account
}

// NewAccountDelivery ...
func NewAccountDelivery(AccountU usecase.Account) AccountDelivery {
	return AccountDelivery{AccountUc: AccountU}
}

// Mount ....
func (h *AccountDelivery) Mount(group *echo.Group, publicKey *rsa.PublicKey) {
	group.GET("", h.ViewAccount)
}

// ViewAccount ...
func (h *AccountDelivery) ViewAccount(c echo.Context) error {
	resp := shared.JSONSuccess(`Success`, nil)
	return c.JSON(resp.Code, resp)
}
