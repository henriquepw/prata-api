package session

import (
	"context"
	"time"

	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/henriquepw/prata-api/pkg/jwt"
)

const ONE_MONTH = time.Hour * 24 * 30

type SessionService interface {
	CreateSession(ctx context.Context, userID id.ID) (Session, error)
	GetByID(ctx context.Context, sessionID id.ID) (Session, error)
}

type sessionService struct {
	store SessionStore
}

func NewService(store SessionStore) SessionService {
	return &sessionService{store}
}

func (s *sessionService) CreateSession(ctx context.Context, userID id.ID) (Session, error) {
	token, claims, err := jwt.Generate(userID.String(), ONE_MONTH)
	if err != nil {
		return Session{}, errorx.Internal()
	}

	session := Session{
		ID:           claims.SessionID,
		ExpiresAt:    claims.ExpiresAt.Time,
		UserID:       userID,
		RefreshToken: token,
	}

	err = s.store.Insert(ctx, session)
	if err != nil {
		return Session{}, err
	}

	return session, err
}

func (s *sessionService) GetByID(ctx context.Context, sessionID id.ID) (Session, error) {
	session, err := s.store.Get(ctx, sessionID)
	if err != nil {
		return Session{}, errorx.NotFound("session not found")
	}

	return session, err
}
