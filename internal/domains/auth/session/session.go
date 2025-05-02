package session

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/jwt"
)

type Session struct {
	ID           id.ID     `json:"id" db:"id"`
	UserID       id.ID     `json:"userId" db:"user_id"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
	ExpiresAt    time.Time `json:"expiresAt" db:"expires_at"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type Access struct {
	SessionID             id.ID     `json:"sessionId"`
	AccessToken           string    `json:"accessToken"`
	RefreshToken          string    `json:"refreshToken"`
	AccessTokenExpiresAt  time.Time `json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
}

func (s *Session) GetAccess() (*Access, error) {
	accessToken, accessClaims, err := jwt.Generate(s.UserID.String(), time.Minute*15)
	if err != nil {
		return nil, errorx.Internal()
	}

	access := Access{
		SessionID:             s.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessClaims.ExpiresAt.Time,
		RefreshToken:          s.RefreshToken,
		RefreshTokenExpiresAt: s.ExpiresAt,
	}
	return &access, nil
}
