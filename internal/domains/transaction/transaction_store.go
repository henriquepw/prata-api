package transaction

import (
	"context"
	"strings"
	"time"

	"github.com/henriquepw/prata-api/pkg/date"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/henriquepw/prata-api/pkg/page"
	"github.com/jmoiron/sqlx"
)

type TransactionStore interface {
	Insert(ctx context.Context, i Transaction) error
	Delete(ctx context.Context, id id.ID) error
	Update(ctx context.Context, id id.ID, i TransactionUpdate) error
	Get(ctx context.Context, id id.ID) (Transaction, error)
	List(ctx context.Context, q TransactionQuery) (page.Cursor[Transaction], error)
}

type transactionStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) TransactionStore {
	return &transactionStore{db}
}

func (s *transactionStore) Insert(ctx context.Context, i Transaction) error {
	query := `
    INSERT INTO transactions (id, user_id, balance_id, type, description, amount, received_at, created_at, updated_at)
    VALUES (:id, :user_id, :balance_id, :type, :description, :amount, :received_at, :created_at, :updated_at)
  `
	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *transactionStore) Delete(ctx context.Context, id id.ID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM transactions WHERE id = ?", id)
	return err
}

func (s *transactionStore) Update(ctx context.Context, id id.ID, i TransactionUpdate) error {
	args := []any{}
	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE transactions set updated_at = ?")
	args = append(args, date.FormatToISO(time.Now()))

	if i.Amount != 0 {
		queryBuilder.WriteString(", amount = ?")
		args = append(args, i.Amount)
	}

	if i.Description != "" {
		queryBuilder.WriteString(", description = ?")
		args = append(args, i.Description)
	}

	if !i.ReceivedAt.IsZero() {
		queryBuilder.WriteString(", received_at = ?")
		args = append(args, date.FormatToISO(i.ReceivedAt))
	}

	queryBuilder.WriteString(" WHERE id = ?")
	args = append(args, id)

	_, err := s.db.ExecContext(ctx, queryBuilder.String(), args...)
	return err
}

func (s *transactionStore) Get(ctx context.Context, id id.ID) (Transaction, error) {
	query := "SELECT * FROM transactions WHERE id = ?"

	var transaction Transaction
	err := s.db.GetContext(ctx, &transaction, query, id)

	return transaction, err
}

func (s *transactionStore) List(ctx context.Context, q TransactionQuery) (page.Cursor[Transaction], error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT * FROM transactions")

	where := []string{}
	args := []any{}

	if q.Cursor != "" {
		where = append(where, "id > ?")
		args = append(args, q.Cursor)
	}

	if q.Search != "" {
		where = append(where, "description LIKE %?%")
		args = append(args, q.Search)
	}

	if !q.ReceivedAtGte.IsZero() {
		where = append(where, "start_at >= ?")
		args = append(args, date.FormatToISO(q.ReceivedAtGte))
	}
	if !q.ReceivedAtLte.IsZero() {
		where = append(where, "start_at <= ?")
		args = append(args, date.FormatToISO(q.ReceivedAtLte))
	}

	if len(where) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(where, " AND "))
	}

	if q.Limit > 0 {
		queryBuilder.WriteString(" Limit ?")
		args = append(args, q.Limit+1)
	}

	var transactions []Transaction
	err := s.db.Select(&transactions, queryBuilder.String(), args...)
	if err != nil {
		return page.NewEmpty[Transaction](), err
	}

	page := page.New(transactions, q.Limit, func(i Transaction) string {
		return i.ID.String()
	})

	return page, nil
}
