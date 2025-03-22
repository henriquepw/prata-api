package recurrence

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type RecurrenceService interface {
	CreateRecurrence(ctx context.Context, dto RecurrenceCreate) (*Recurrence, error)
	UpdateRecurrence(ctx context.Context, id id.ID, dto RecurrenceUpdate) error
	DeleteRecurrence(ctx context.Context, id id.ID) error
	GetRecurrence(ctx context.Context, id id.ID) (*Recurrence, error)
	ListRecurrence(ctx context.Context, dto RecurrenceQuery) *page.Cursor[Recurrence]
}

type recurrenceService struct {
	store RecurrenceStore
}

func NewService(store RecurrenceStore) RecurrenceService {
	return &recurrenceService{store}
}

func (s *recurrenceService) CreateRecurrence(ctx context.Context, dto RecurrenceCreate) (*Recurrence, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	now := time.Now()
	recurrence := Recurrence{
		ID:           id.New(),
		AccountID:    dto.AccountID,
		Description:  dto.Description,
		Frequence:    dto.Frequence,
		Installments: dto.Installments,
		StartAt:      dto.StartAt,
		EndAt:        dto.EndAt,
		CreatedAt:    now,
		UpdatedAt:    now,
		DeletedAt:    nil,
	}

	err := s.store.Insert(ctx, recurrence)
	if err != nil {
		return nil, errors.Internal("Failed to create the recurrence")
	}

	return &recurrence, nil
}

func (s *recurrenceService) UpdateRecurrence(ctx context.Context, id id.ID, dto RecurrenceUpdate) error {
	if err := validate.Check(dto); err != nil {
		return err
	}

	// TODO:

	return nil
}

func (s *recurrenceService) DeleteRecurrence(ctx context.Context, id id.ID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Internal("Failed to delete the recurrence")
	}

	return nil
}

func (s *recurrenceService) GetRecurrence(ctx context.Context, id id.ID) (*Recurrence, error) {
	recurrence, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Recurrence not found")
	}

	return recurrence, nil
}

func (s *recurrenceService) ListRecurrence(ctx context.Context, dto RecurrenceQuery) *page.Cursor[Recurrence] {
	if err := validate.Check(dto); err != nil {
		return page.NewEmpty[Recurrence]()
	}

	items, err := s.store.List(ctx, dto)
	if err != nil {
		return page.NewEmpty[Recurrence]()
	}

	return items
}
