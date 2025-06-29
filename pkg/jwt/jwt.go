package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henriquepw/prata-api/internal/env"
	"github.com/henriquepw/prata-api/pkg/id"
)

type Claims struct {
	SessionID id.ID
	jwt.RegisteredClaims
}

func Generate(subject string, duration time.Duration) (string, *Claims, error) {
	now := time.Now()
	id := id.New()
	claims := Claims{
		SessionID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        id.String(),
			Subject:   subject,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(os.Getenv(env.JWTSecret)))

	return token, &claims, err
}

func Validade(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv(env.JWTSecret)), nil
	})
	if err != nil {
		return nil, errors.New("can't parse token")
	}

	claims, ok := t.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	return claims, err
}
