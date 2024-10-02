package middleware

import (
	"github.com/banggibima/backend-agile/config"
)

type JWTMiddleware struct {
	Config config.Config
}

func NewJWTMiddleware(config config.Config) JWTMiddleware {
	return JWTMiddleware{
		Config: config,
	}
}
