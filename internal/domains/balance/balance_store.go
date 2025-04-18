package balance

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type BalanceStore interface {
	Upsert(ctx context.Context, i Balance) error
	Get(ctx context.Context, userID string) (Balance, error)
}

type balanceStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) BalanceStore {
	return &balanceStore{db}
}

func (s *balanceStore) Upsert(ctx context.Context, balance Balance) error {
	query := `
    INSERT INTO balances (id, user_id, label, percent, created_at, updated_at)
    VALUES (:id, :user_id, :label, :percent, :created_at, :updated_at)
    ON CONFLICT (id) DO UPDATE
	  SET label = :label, percent = :percent, updated_at = :updated_at
  `

	_, err := s.db.NamedExecContext(ctx, query, balance)
	return err
}

func (s *balanceStore) Get(ctx context.Context, userID string) (Balance, error) {
	query := `SELECT * FROM balances WHERE user_id = ?`

	var balance Balance
	err := s.db.SelectContext(ctx, &balance, query, userID)

	return balance, err
}
