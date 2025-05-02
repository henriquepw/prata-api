package auth

import (
	"time"
)

type SignInData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RenewAccess struct {
	AccesToken           string    `json:"accessToken"`
	AccessTokenExpiresAt time.Time `json:"accessTokenExpiresAt"`
}
