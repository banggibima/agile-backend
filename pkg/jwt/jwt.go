package jwt

import (
	"time"

	"github.com/banggibima/agile-backend/internal/module/user/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	Secret   string
	Expire   int
	Audience string
	Issuer   string
}

type CustomClaims struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
	jwt.RegisteredClaims
}

func Encoded(j *JWT, user *domain.User) (*jwt.Token, error) {
	claims := CustomClaims{
		ID:   user.ID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.Expire))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	raw, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}

func Decoded(j *JWT, raw string) (*jwt.Token, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}
