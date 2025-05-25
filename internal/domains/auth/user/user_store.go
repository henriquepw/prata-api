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
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id id.ID) (User, error)
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
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)`
	err := s.db.GetContext(ctx, &exists, query, email)

	return exists, err
}

func (s *userStore) Insert(ctx context.Context, i User) error {
	query := `
    INSERT INTO users (id, username, email, avatar, secret)
		VALUES (:id, :username, :email, :avatar, :secret)`
	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *userStore) Delete(ctx context.Context, id id.ID) error {
	query := `UPDATE users set deleted_at = ?`
	_, err := s.db.NamedExecContext(ctx, query, date.FormatToISO(time.Now()))

	return err
}

func (s *userStore) GetByID(ctx context.Context, id id.ID) (User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	var user User
	err := s.db.GetContext(ctx, &user, query, id)

	return user, err
}

func (s *userStore) GetByEmail(ctx context.Context, email string) (User, error) {
	query := "SELECT * FROM users WHERE email = ?"

	var user User
	err := s.db.GetContext(ctx, &user, query, email)

	return user, err
}
