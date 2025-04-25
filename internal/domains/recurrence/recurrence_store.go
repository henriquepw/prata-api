package recurrence

import (
	"context"
	"strings"
	"time"

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
	TodayRecurrences(ctx context.Context) []Recurrence
}

type recurrenceStore struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) RecurrenceStore {
	return &recurrenceStore{db}
}

func (s *recurrenceStore) Insert(ctx context.Context, i Recurrence) error {
	query := `
    INSERT INTO recurrences (id, user_id, description, frequence, day, week, month, year_day, start_at, end_at, created_at, updated_at)
		VALUES (:id, :user_id, :description, :frequence, :day, :week, :month, :year_day, :start_at, :end_at, :created_at, :updated_at)
  `

	_, err := s.db.NamedExecContext(ctx, query, i)
	return err
}

func (s *recurrenceStore) Delete(ctx context.Context, id id.ID) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM recurrences WHERE id = ?", id)
	return err
}

func (s *recurrenceStore) Update(ctx context.Context, id id.ID, i RecurrenceUpdate) error {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("UPDATE recurrences SET updated_at = ?")
	args := []any{date.FormatToISO(time.Now())}

	if i.Amount != 0 {
		queryBuilder.WriteString(", amount = ?")
		args = append(args, i.Amount)
	}

	if i.Description != "" {
		queryBuilder.WriteString(", description = ?")
		args = append(args, i.Description)
	}

	if i.Frequence != "" {
		queryBuilder.WriteString(", frequence = ?")
		args = append(args, i.Frequence)
	}

	if i.BalanceID == nil || *i.BalanceID != "" {
		queryBuilder.WriteString(", balance_id = ?")
		args = append(args, i.BalanceID)
	}

	if i.EndAt == nil || !i.EndAt.IsZero() {
		queryBuilder.WriteString(", end_at = ?")
		args = append(args, date.FormatToISO(*i.EndAt))
	}

	queryBuilder.WriteString(" WHERE id ?")
	args = append(args, id)

	_, err := s.db.ExecContext(ctx, queryBuilder.String(), args)
	return err
}

func (s *recurrenceStore) Get(ctx context.Context, id id.ID) (*Recurrence, error) {
	query := `SELECT * FROM recurrences WHERE id = ?`

	var recurrence Recurrence
	err := s.db.GetContext(ctx, &recurrence, query, id)
	if err != nil {
		return nil, err
	}

	return &recurrence, nil
}

func (s *recurrenceStore) List(ctx context.Context, q RecurrenceQuery) (*page.Cursor[Recurrence], error) {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT * FROM recurrences")

	where := []string{}
	args := []any{}

	if q.Cursor != "" {
		where = append(where, "id > ?")
		args = append(args, q.Cursor)
	}
	if q.Frequence != "" {
		where = append(where, "frequence = ?")
		args = append(args, q.Frequence)
	}
	if q.Search != "" {
		where = append(where, "description LIKE %?%")
		args = append(args, q.Search)
	}

	if !q.StartAtGte.IsZero() {
		where = append(where, "start_at >= ?")
		args = append(args, date.FormatToISO(q.StartAtGte))
	}
	if !q.StartAtLte.IsZero() {
		where = append(where, "start_at <= ?")
		args = append(args, date.FormatToISO(q.StartAtLte))
	}

	if !q.EndAtGte.IsZero() {
		where = append(where, "end_at >= ?")
		args = append(args, date.FormatToISO(q.EndAtGte))
	}
	if !q.EndAtLte.IsZero() {
		where = append(where, "end_at <= ?")
		args = append(args, date.FormatToISO(q.EndAtLte))
	}

	if len(where) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(where, " AND "))
	}

	if q.Limit > 0 {
		queryBuilder.WriteString(" Limit ?")
		args = append(args, q.Limit+1)
	}

	var recurrences []Recurrence
	err := s.db.Select(&recurrences, queryBuilder.String(), args...)
	if err != nil {
		return nil, err
	}

	page := page.New(recurrences, q.Limit, func(i Recurrence) string {
		return i.ID.String()
	})

	return page, nil
}

func (s *recurrenceStore) TodayRecurrences(ctx context.Context) []Recurrence {
	now := time.Now()
	nextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := nextMonth.AddDate(0, 0, -1).Day()
	filters := map[string]any{
		"day":            now.Day(),      // 1-31
		"week":           now.Weekday(),  // 0-6
		"month":          now.Month(),    // 1-12
		"yearDay":        now.YearDay(),  // 1-366
		"lastDayOfMonth": lastDayOfMonth, // 28-31
	}

	query := `
		SELECT *
		FROM recurrences
		WHERE (end_at IS NULL OR end_at > DATE('now'))
			AND (
				(frequence = 'WEEKLY' AND week = :week)
	      OR (frequence = 'BIWEEKLY' AND :yearDay - year_day % 14 = 0)
				OR (
					frequence = 'MONTHLY'
					AND (
						day = :day
						OR (:day = :lastDayOfMonth AND day > :lastDayOfMonth)
					)
				)
				OR (frequence = 'YEARLY' AND day = :day AND month = :month)
			)
	`

	var recurrences []Recurrence
	s.db.Select(&recurrences, query, filters)

	return recurrences
}
