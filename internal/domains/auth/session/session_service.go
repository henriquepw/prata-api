package session

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/jwt"
)

type SessionService interface {
	CreateSession(ctx context.Context, userID id.ID) (*Session, error)
	GetByID(ctx context.Context, sessionID id.ID) (*Session, error)
}

type sessionService struct {
	store SessionStore
}

func NewService(store SessionStore) SessionService {
	return &sessionService{store}
}

func (s *sessionService) CreateSession(ctx context.Context, userID id.ID) (*Session, error) {
	refreshToken, refreshClaims, err := jwt.Generate(userID.String(), time.Hour*24*30)
	if err != nil {
		return nil, errorx.Internal()
	}

	session := Session{
		ID:           refreshClaims.ID,
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshClaims.ExpiresAt.Time,
	}

	err = s.store.Insert(ctx, session)
	if err != nil {
		return nil, err
	}

	return &session, err
}

func (s *sessionService) GetByID(ctx context.Context, sessionID id.ID) (*Session, error) {
	session, err := s.store.Get(ctx, sessionID)
	if err != nil {
		return nil, errorx.NotFound("user not found")
	}

	return session, err
}
