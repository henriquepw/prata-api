package user

import (
	"time"

	"github.com/henriquepw/prata-api/pkg/id"
)

type User struct {
	ID        id.ID     `db:"id"`
	Email     string     `db:"email"`
	Secret    string     `db:"secret"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
