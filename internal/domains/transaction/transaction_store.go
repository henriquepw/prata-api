package transaction

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/date"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/jmoiron/sqlx"
)

type TransactionStore interface {
	Insert(ctx context.Context, i Transaction) error
	Delete(ctx context.Context, id id.ID) error
	Update(ctx context.Context, id id.ID, i TransactionUpdate) error
	Get(ctx context.Context, id id.ID) (*Transaction, error)
	List(ctx context.Context, q TransactionQuery) (*page.Cursor[Transaction], error)
}

type transactionStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) TransactionStore {
	return &transactionStore{db}
}

func (s *transactionStore) Insert(ctx context.Context, i Transaction) error {
	query := `
    INSERT INTO transactions (id, account_id, tags, type, description, amount, received_at, created_at, updated_at)
    VALUES (:id, :account_id, :tags, :type, :description, :amount, :received_at, :created_at, :updated_at)
  `
	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *transactionStore) Delete(ctx context.Context, id id.ID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM transactions WHERE id = ?", id)
	return err
}

func (s *transactionStore) Update(ctx context.Context, id id.ID, i TransactionUpdate) error {
	query := `
    UPDATE transactions
    SET amount = ?, received_at = ?, updated_at = ?
    WHERE id = ?
  `
	_, err := s.db.ExecContext(
		ctx, query,
		i.Amount,
		date.FormatToISO(i.ReceivedAt),
		date.FormatToISO(time.Now()),
		id,
	)

	return err
}

func (s *transactionStore) Get(ctx context.Context, id id.ID) (*Transaction, error) {
	query := "SELECT * FROM transactions WHERE id = ?"

	var transaction Transaction
	err := s.db.GetContext(ctx, &transaction, query, id)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *transactionStore) List(ctx context.Context, q TransactionQuery) (*page.Cursor[Transaction], error) {
	query := `
    SELECT *
    FROM transactions
    WHERE received_at > ?
    ORDER BY received_at ASC
    LIMIT ?
  `

	var transactions []Transaction
	err := s.db.Select(&transactions, query, q.Cursor, q.Limit+1)
	if err != nil {
		return nil, err
	}

	page := page.New(transactions, q.Limit, func(i Transaction) string {
		return date.FormatToISO(i.ReceivedAt)
	})

	return page, nil
}
