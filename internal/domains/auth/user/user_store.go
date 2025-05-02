package user

import (
	"context"
	"time"

	"github.com/henriquepw/prata-api/pkg/date"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	Insert(ctx context.Context, i User) error
	Delete(ctx context.Context, id id.ID) error
	Get(ctx context.Context, email string) (*User, error)
	Has(ctx context.Context, email string) (bool, error)
}

type userStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) UserStore {
	return &userStore{db}
}

func (s *userStore) Has(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = ?);`
	err := s.db.Get(ctx, query, email)

	return exists, err
}

func (s *userStore) Insert(ctx context.Context, i User) error {
	query := `
    INSERT INTO users (id, email, secret)
		VALUES (:id, :email, :secret)`
	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *userStore) Delete(ctx context.Context, id id.ID) error {
	query := `UPDATE users set deleted_at = ?`
	_, err := s.db.NamedExecContext(ctx, query, date.FormatToISO(time.Now()))

	return err
}

func (s *userStore) Get(ctx context.Context, email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"

	var user User
	err := s.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
