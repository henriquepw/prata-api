package auth

import (
	"context"

	"github.com/henriquepw/prata-api/internal/domains/auth/session"
	"github.com/henriquepw/prata-api/internal/domains/auth/user"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/hash"
	"github.com/henriquepw/prata-api/pkg/jwt"
	"github.com/henriquepw/prata-api/pkg/validate"
)

type AuthService interface {
	SignIn(ctx context.Context, dto SignInData) (*session.Access, error)
	SignUp(ctx context.Context, dto SignUpData) (*session.Access, error)
	RefreshAccess(ctx context.Context, refreshToken string) (*RenewAccess, error)
}

type authService struct {
	userSVC    user.UserService
	sessionSVC session.SessionService
}

func NewService(userSVC user.UserService, sessionSVC session.SessionService) AuthService {
	return &authService{userSVC, sessionSVC}
}

func (s *authService) SignUp(ctx context.Context, dto SignUpData) (*session.Access, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	// create user
	user, err := s.userSVC.CreateUser(ctx, dto.Email, dto.Password)
	if err != nil {
		return nil, err
	}

	// create session
	session, err := s.sessionSVC.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	access, err := session.GetAccess()
	if err != nil {
		return nil, errorx.Internal()
	}

	return access, nil
}

func (s *authService) SignIn(ctx context.Context, dto SignInData) (*session.Access, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	// get user and validate password
	user, err := s.userSVC.GetByEmail(ctx, dto.Email)
	if err != nil {
		return nil, err
	}

	if !hash.Validate(user.Secret, dto.Password) {
		return nil, errorx.Unauthorized()
	}

	// create session
	session, err := s.sessionSVC.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	access, err := session.GetAccess()
	if err != nil {
		return nil, errorx.Internal()
	}

	return access, nil
}

func (s *authService) RefreshAccess(ctx context.Context, refreshToken string) (*RenewAccess, error) {
	claims, err := jwt.Validade(refreshToken)
	if err != nil {
		return nil, errorx.Unauthorized()
	}

	session, err := s.sessionSVC.GetByID(ctx, claims.ID)
	if err != nil {
		return nil, errorx.Unauthorized()
	}

	if session.UserID != claims.ID {
		return nil, errorx.Unauthorized()
	}

	access, err := session.GetAccess()
	if err != nil {
		return nil, errorx.Internal()
	}

	return &RenewAccess{
		AccesToken:           access.AccessToken,
		AccessTokenExpiresAt: access.AccessTokenExpiresAt,
	}, nil
}
