package recurrence

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/henriquepw/prata-api/pkg/page"
	"github.com/henriquepw/prata-api/pkg/validate"
)

type RecurrenceService interface {
	CreateRecurrence(ctx context.Context, dto RecurrenceCreate) (Recurrence, error)
	UpdateRecurrence(ctx context.Context, id id.ID, dto RecurrenceUpdate) error
	DeleteRecurrence(ctx context.Context, id id.ID) error
	GetRecurrence(ctx context.Context, id id.ID) (Recurrence, error)
	ListRecurrence(ctx context.Context, dto RecurrenceQuery) page.Cursor[Recurrence]
}

type recurrenceService struct {
	store RecurrenceStore
}

func NewService(store RecurrenceStore) RecurrenceService {
	return &recurrenceService{store}
}

func (s *recurrenceService) CreateRecurrence(ctx context.Context, dto RecurrenceCreate) (Recurrence, error) {
	if err := validate.Check(dto); err != nil {
		return Recurrence{}, err
	}

	now := time.Now()
	recurrence := Recurrence{
		ID:          id.New(),
		UserID:      dto.UserID,
		Amount:      dto.Amount,
		Description: dto.Description,
		Type:        dto.Type,
		Frequence:   dto.Frequence,
		StartAt:     dto.StartAt,
		Day:         dto.StartAt.Day(),
		Week:        int(dto.StartAt.Weekday()),
		Month:       int(dto.StartAt.Month()),
		YearDay:     dto.StartAt.YearDay(),
		EndAt:       dto.EndAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.store.Insert(ctx, recurrence)
	if err != nil {
		log.Error("Failed to create the recurrence", err)
		return Recurrence{}, errorx.Internal("Failed to create the recurrence")
	}

	return recurrence, nil
}

func (s *recurrenceService) UpdateRecurrence(ctx context.Context, id id.ID, dto RecurrenceUpdate) error {
	if err := validate.Check(dto); err != nil {
		return err
	}

	err := s.store.Update(ctx, id, dto)
	if err != nil {
		return errorx.Internal("Can't update the recurrence")
	}

	return nil
}

func (s *recurrenceService) DeleteRecurrence(ctx context.Context, id id.ID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errorx.Internal("Failed to delete the recurrence")
	}

	return nil
}

func (s *recurrenceService) GetRecurrence(ctx context.Context, id id.ID) (Recurrence, error) {
	recurrence, err := s.store.Get(ctx, id)
	if err != nil {
		return recurrence, errorx.NotFound("Recurrence not found")
	}

	return recurrence, nil
}

func (s *recurrenceService) ListRecurrence(ctx context.Context, dto RecurrenceQuery) page.Cursor[Recurrence] {
	if err := validate.Check(dto); err != nil {
		return page.NewEmpty[Recurrence]()
	}

	items, err := s.store.List(ctx, dto)
	if err != nil {
		return page.NewEmpty[Recurrence]()
	}

	return items
}
