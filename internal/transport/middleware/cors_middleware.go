package middleware

import (
	"github.com/banggibima/agile-backend/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CORSMiddleware struct {
	Config *config.Config
}

func NewCORSMiddleware(config *config.Config) CORSMiddleware {
	return CORSMiddleware{
		Config: config,
	}
}

func (c CORSMiddleware) WithConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(c.BuildCORSConfig())
}

func (c CORSMiddleware) BuildCORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: c.GetAllowedOrigins(),
		AllowHeaders: c.GetAllowedHeaders(),
	}
}

func (c CORSMiddleware) GetAllowedOrigins() []string {
	if len(c.Config.App.Env) > 0 && c.Config.App.Env == "production" {
		return []string{"https://example.com"}
	}
	return []string{"*"}
}

func (c CORSMiddleware) GetAllowedHeaders() []string {
	return []string{
		echo.HeaderOrigin,
		echo.HeaderContentType,
		echo.HeaderAccept,
	}
}
