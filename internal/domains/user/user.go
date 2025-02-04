package user

import (
	"database/sql"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
)

type User struct {
	ID        id.ID        `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deletedAt" db:"deleted_at"`
	LastLogin sql.NullTime `json:"lastLogin" db:"last_login"`

	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
