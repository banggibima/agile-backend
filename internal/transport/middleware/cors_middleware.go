package middleware

import (
	"github.com/banggibima/backend-agile/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSMiddleware struct {
	Config *config.Config
}

func NewCORSMiddleware(
	config *config.Config,
) CORSMiddleware {
	return CORSMiddleware{
		Config: config,
	}
}

func (c CORSMiddleware) WithConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}
