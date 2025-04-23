package auth

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/henriquepw/pobrin-api/pkg/errorx"
)

type Session interface {
	GetUserID(ctx context.Context) (string, error)
}

type session struct{}

func NewSession() Session {
	return session{}
}

func (s session) GetUserID(ctx context.Context) (string, error) {
	claims, ok := clerk.SessionClaimsFromContext(ctx)
	if !ok {
		return "", errorx.Unauthorized()
	}

	return claims.Subject, nil
}
