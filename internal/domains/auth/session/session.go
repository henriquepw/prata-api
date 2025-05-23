package session

import (
	"os"
	"strconv"
	"time"

	"github.com/henriquepw/prata-api/internal/env"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/henriquepw/prata-api/pkg/jwt"
)

type Session struct {
	ID           id.ID     `json:"id" db:"id"`
	UserID       id.ID     `json:"userId" db:"user_id"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
	ExpiresAt    time.Time `json:"expiresAt" db:"expires_at"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

type Access struct {
	UserID                id.ID     `json:"userId"`
	AccessToken           string    `json:"accessToken"`
	RefreshToken          string    `json:"refreshToken"`
	AccessTokenExpiresAt  time.Time `json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
}

func (s Session) GetAccess() (*Access, error) {
	accessTime, err := strconv.Atoi(os.Getenv(env.ACCESS_TIME))
	if err != nil {
		return nil, errorx.Internal("invalid acesss time")
	}

	token, claims, err := jwt.Generate(s.UserID.String(), time.Minute*time.Duration(accessTime))
	if err != nil {
		return nil, errorx.Internal("can't generate the access token")
	}

	access := Access{
		UserID:                s.UserID,
		AccessToken:           token,
		AccessTokenExpiresAt:  claims.ExpiresAt.Time,
		RefreshToken:          s.RefreshToken,
		RefreshTokenExpiresAt: s.ExpiresAt,
	}
	return &access, nil
}
