// Package auth implements the user authorization workflow
package auth

import (
	"time"
)

type SignInData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6"`
}

type SignUpData struct {
	Avatar   string `json:"avatar"  validate:"omitempty"`
	Username string `json:"username"  validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=6"`
}

type RewewData struct {
	RefreshToken string `json:"refreshToken"`
}

type RenewAccess struct {
	AccesToken string    `json:"accessToken"`
	ExpiresAt  time.Time `json:"expiresAt"`
}
