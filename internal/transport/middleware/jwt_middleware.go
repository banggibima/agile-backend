package middleware

import (
	"github.com/banggibima/agile-backend/config"
)

type JWTMiddleware struct {
	Config config.Config
}

func NewJWTMiddleware(config config.Config) JWTMiddleware {
	return JWTMiddleware{
		Config: config,
	}
}
