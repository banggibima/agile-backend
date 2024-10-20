package auth

import (
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	Logout(c echo.Context) error
}

type AuthRouter interface {
	Resource() error
}
