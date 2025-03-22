package recurrence

import (
	"context"

	"github.com/henriquepw/pobrin-api/pkg/date"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/jmoiron/sqlx"
)

type RecurrenceStore interface {
	Insert(ctx context.Context, i Recurrence) error
	Delete(ctx context.Context, id id.ID) error
	Update(ctx context.Context, id id.ID, i RecurrenceUpdate) error
	Get(ctx context.Context, id id.ID) (*Recurrence, error)
	List(ctx context.Context, q RecurrenceQuery) (*page.Cursor[Recurrence], error)
}

type recurrenceStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) RecurrenceStore {
	return &recurrenceStore{db}
}

func (s *recurrenceStore) Insert(ctx context.Context, i Recurrence) error {
	query := `
    INSERT INTO recurrence (id, description, frequence, installments, start_at, end_at, created_at, updated_at)
    VALUES (:id, :description, :frequence, :installments, :start_at, :end_at, :created_at, :updated_at)
  `

	_, err := s.db.NamedExecContext(ctx, query, i)

	return err
}

func (s *recurrenceStore) Delete(ctx context.Context, id id.ID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM recurrence WHERE id = ?", id)
	return err
}

func (s *recurrenceStore) Update(ctx context.Context, id id.ID, i RecurrenceUpdate) error {
	return nil
	// query := `
	//    UPDATE recurrence
	//    SET amount = ?, received_at = ?, updated_at = ?
	//    WHERE id = ?
	//  `
	// _, err := s.db.ExecContext(
	// 	ctx, query,
	// 	i.Amount,
	// 	date.FormatToISO(i.ReceivedAt),
	// 	date.FormatToISO(time.Now()),
	// 	id,
	// )
	//
	// return err
}

func (s *recurrenceStore) Get(ctx context.Context, id id.ID) (*Recurrence, error) {
	query := `
    SELECT id, description, frequence, installments, start_at, end_at, created_at, updated_at, deleted_at
    FROM recurrence
    WHERE id = ?
  `

	var recurrence Recurrence
	err := s.db.GetContext(ctx, &recurrence, query, id)
	if err != nil {
		return nil, err
	}

	return &recurrence, nil
}

func (s *recurrenceStore) List(ctx context.Context, q RecurrenceQuery) (*page.Cursor[Recurrence], error) {
	query := `
    SELECT id, description, frequence, installments, start_at, end_at, created_at, updated_at, deleted_at
    FROM recurrence
    WHERE created_at > ?
    ORDER BY created_at ASC
    LIMIT ?
  `

	var recurrences []Recurrence
	err := s.db.Select(&recurrences, query, q.Cursor, q.Limit+1)
	if err != nil {
		return nil, err
	}

	page := page.New(recurrences, q.Limit, func(i Recurrence) string {
		return date.FormatToISO(i.CreatedAt)
	})

	return page, nil
}
