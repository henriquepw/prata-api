package balance

import (
	"context"

	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/jmoiron/sqlx"
)

type BalanceStore interface {
	Upsert(ctx context.Context, i Balance) error
	Get(ctx context.Context, userID id.ID) (Balance, error)
}

type balanceStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) BalanceStore {
	return &balanceStore{db}
}

func (s *balanceStore) Upsert(ctx context.Context, balance Balance) error {
	query := `
    INSERT INTO balances (id, user_id, label, color, percent, created_at, updated_at)
	  VALUES (:id, :user_id, :label, :color, :percent, :created_at, :updated_at)
    ON CONFLICT (id) DO UPDATE SET label=excluded.label, color=excluded.color, percent=excluded.percent, updated_at=excluded.updated_at
  `

	_, err := s.db.NamedExecContext(ctx, query, balance.Pieces)
	return err
}

func (s *balanceStore) Get(ctx context.Context, userID id.ID) (Balance, error) {
	query := `SELECT * FROM balances WHERE user_id = ?`

	balance := Balance{Pieces: []Piece{}}
	err := s.db.SelectContext(ctx, &balance.Pieces, query, userID)

	return balance, err
}
