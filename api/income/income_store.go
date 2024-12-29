package income

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/date"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/jmoiron/sqlx"
)

type IncomeStore interface {
	Insert(ctx context.Context, i Income) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, i IncomeUpdate) error
	Get(ctx context.Context, id string) (*Income, error)
	List(ctx context.Context, q IncomeQuery) (*page.Cursor[Income], error)
}

type incomeStore struct {
	db *sqlx.DB
}

func NewIncomeStore(db *sqlx.DB) IncomeStore {
	return &incomeStore{db}
}

func (s *incomeStore) Insert(ctx context.Context, i Income) error {
	query := `
    INSERT INTO Income (id, amount, received_at, created_at, updated_at)
    VALUES (:id, :amount, :received_at, :created_at, :updated_at)
  `
	_, error := s.db.NamedExecContext(ctx, query, i)

	return error
}

func (s *incomeStore) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM income WHERE id = ?", id)
	return err
}

func (s *incomeStore) Update(ctx context.Context, id string, i IncomeUpdate) error {
	query := `
    UPDATE income
    SET amount = ?, received_at = ?, updated_at = ?
    WHERE id = ?
  `
	_, error := s.db.ExecContext(
		ctx, query,
		i.Amount,
		date.FormatToISO(i.ReceivedAt),
		date.FormatToISO(time.Now()),
		id,
	)

	return error
}

func (s *incomeStore) Get(ctx context.Context, id string) (*Income, error) {
	query := `
    SELECT id, amount, received_at, created_at, updated_at
    FROM Income
    WHERE id = ?
  `

	var income Income
	err := s.db.GetContext(ctx, &income, query, id)
	if err != nil {
		return nil, err
	}

	return &income, nil
}

func (s *incomeStore) List(ctx context.Context, q IncomeQuery) (*page.Cursor[Income], error) {
	var incomes []Income

	query := `
    SELECT id, amount, received_at, created_at, updated_at
    FROM Income
    WHERE received_at > ?
    ORDER BY received_at ASC
    LIMIT ?
  `
	err := s.db.Select(&incomes, query, q.Cursor, q.Limit+1)
	if err != nil {
		return nil, err
	}

	page := page.New(incomes, q.Limit, func(i Income) string {
		return date.FormatToISO(i.ReceivedAt)
	})

	return page, nil
}
