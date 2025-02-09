package jwtutils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henriquepw/pobrin-api/pkg/id"
)

const (
	accessTokenTimeout  = 15 * time.Minute
	refreshTokenTimeout = 24 * 30 * time.Hour
	mySigningKey        = "AllYourBase"
)

var claims = jwt.MapClaims{}

type Claims struct {
	jwt.RegisteredClaims
	TokenType string `json:"tokenType"`
}

func NewClaims(userID id.ID, t string) *Claims {
	now := time.Now()
	c := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(accessTokenTimeout)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    "https://api.pobrin.com.br",
		Subject:   userID.String(),
		ID:        id.NewTiny().String(),
		Audience:  []string{"pobrin-api"},
	}

	return &Claims{
		RegisteredClaims: c,
		TokenType:        t,
	}
}
