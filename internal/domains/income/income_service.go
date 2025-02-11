package income

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type IncomeService interface {
	CreateIncome(ctx context.Context, dto IncomeCreate) (*Income, error)
	UpdateIncome(ctx context.Context, id id.ID, dto IncomeUpdate) error
	DeleteIncome(ctx context.Context, id id.ID) error
	GetIncome(ctx context.Context, id id.ID) (*Income, error)
	ListIncome(ctx context.Context, dto IncomeQuery) *page.Cursor[Income]
}

type incomeService struct {
	store IncomeStore
}

func NewIncomeService(store IncomeStore) IncomeService {
	return &incomeService{store}
}

func (s *incomeService) CreateIncome(ctx context.Context, dto IncomeCreate) (*Income, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	now := time.Now()
	income := Income{
		ID:         id.New(),
		Amount:     dto.Amount,
		ReceivedAt: dto.ReceivedAt,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err := s.store.Insert(ctx, income)
	if err != nil {
		return nil, errors.Internal("Failed to create the income")
	}

	return &income, nil
}

func (s *incomeService) UpdateIncome(ctx context.Context, id id.ID, dto IncomeUpdate) error {
	if err := validate.Check(dto); err != nil {
		return err
	}

	return nil
}

func (s *incomeService) DeleteIncome(ctx context.Context, id id.ID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Internal("Failed to delete the income")
	}

	return nil
}

func (s *incomeService) GetIncome(ctx context.Context, id id.ID) (*Income, error) {
	income, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Income not found")
	}

	return income, nil
}

func (s *incomeService) ListIncome(ctx context.Context, dto IncomeQuery) *page.Cursor[Income] {
	if err := validate.Check(dto); err != nil {
		return page.NewEmpty[Income]()
	}

	items, err := s.store.List(ctx, dto)
	if err != nil {
		return page.NewEmpty[Income]()
	}

	return items
}
