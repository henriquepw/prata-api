package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/henriquepw/pobrin-api/internal/domains/auth/jwtutils"
	"github.com/henriquepw/pobrin-api/internal/domains/user"
	"github.com/henriquepw/pobrin-api/internal/env"
	"github.com/henriquepw/pobrin-api/pkg/hash"
)

type Service interface {
	Login(ctx context.Context, data *LoginRequest) (*LoginResponse, error)
}

type service struct {
	store user.Store
}

func NewService(s user.Store) *service {
	return &service{store: s}
}

func (svc *service) Login(ctx context.Context, data *LoginRequest) (*LoginResponse, error) {
	user, err := svc.store.GetUserPassword(ctx, data.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get stored user password: %w", err)
	}

	if !hash.Validate(user.Password, data.Password) {
		return nil, ErrInvalidPassword
	}

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtutils.NewClaims(user.ID, "access"))
	a, err := access.SignedString([]byte(os.Getenv(env.JWTSecret)))
	if err != nil {
		return nil, err
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtutils.NewClaims(user.ID, "refresh"))
	r, err := refresh.SignedString([]byte(os.Getenv(env.JWTSecret)))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Access: a, Refresh: r}, nil
}
