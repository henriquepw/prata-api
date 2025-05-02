package session

import (
	"context"

	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/jmoiron/sqlx"
)

type SessionStore interface {
	Insert(ctx context.Context, i Session) error
	Delete(ctx context.Context, id id.ID) error
	Get(ctx context.Context, id id.ID) (*Session, error)
}

type sessioStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) SessionStore {
	return &sessioStore{db}
}

func (s *sessioStore) Insert(ctx context.Context, i Session) error {
	query := `
    INSERT INTO sessions (id, user_id, refresh_token)
		VALUES (:id, :user_id, :refresh_token)
	`
	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *sessioStore) Delete(ctx context.Context, id id.ID) error {
	query := `DELETE FROM sessions WHERE id = ?`
	_, err := s.db.NamedExecContext(ctx, query, id)

	return err
}

func (s *sessioStore) Get(ctx context.Context, id id.ID) (*Session, error) {
	query := "SELECT * FROM sessions WHERE id = ?"

	var item Session
	err := s.db.GetContext(ctx, &item, query, id)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
