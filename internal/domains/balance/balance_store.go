package balance

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/date"
	"github.com/jmoiron/sqlx"
)

type BalanceStore interface {
	Insert(ctx context.Context, i Balance) error
	Get(ctx context.Context, userID string) (*Balance, error)
	Update(ctx context.Context, dto BalanceUpdate) error
}

type balanceStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) BalanceStore {
	return &balanceStore{db}
}

func (s *balanceStore) Insert(ctx context.Context, i Balance) error {
	query := `
    INSERT INTO balances (id, user_id, basic, fun, saves, created_at, updated_at)
    VALUES (:id, :user_id, :fun, :saves, :created_at, :updated_at)
  `

	_, err := s.db.NamedExecContext(ctx, query, i)
	return err
}

func (s *balanceStore) Update(ctx context.Context, dto BalanceUpdate) error {
	query := `
	   UPDATE balances
	   SET basic = ?, fun = ?, saves = ?, updated_at = ?
	   WHERE user_id = ?
	 `
	_, err := s.db.ExecContext(
		ctx, query,
		dto.Basic,
		dto.Fun,
		dto.Saves,
		date.FormatToISO(time.Now()),
		dto.UserID,
	)

	return err
}

func (s *balanceStore) Get(ctx context.Context, userId string) (*Balance, error) {
	query := `SELECT * FROM balances WHERE user_id = ?`

	var balance Balance
	err := s.db.GetContext(ctx, &balance, query, userId)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}
