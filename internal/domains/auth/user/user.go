// Package user implements user CRUD
package user

import (
	"time"

	"github.com/henriquepw/prata-api/pkg/id"
)

type User struct {
	ID        id.ID      `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	Username  string     `json:"username" db:"username"`
	Avatar    string     `json:"avatar" db:"avatar"`
	Secret    string     `json:"-" db:"secret"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

type UserCreate struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Avatar   string `json:"avatar" validate:"omitempty"`
	Password string `json:"password" validate:"required"`
}
