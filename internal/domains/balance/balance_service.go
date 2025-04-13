package balance

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type BalanceService interface {
	CreateBalance(ctx context.Context, dto BalanceCreate) (*Balance, error)
	UpdateBalance(ctx context.Context, dto BalanceUpdate) error
	GetBalance(ctx context.Context, userID string) (*Balance, error)
}

type balanceService struct {
	store BalanceStore
}

func NewService(store BalanceStore) BalanceService {
	return &balanceService{store}
}

func (s *balanceService) CreateBalance(ctx context.Context, dto BalanceCreate) (*Balance, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	now := time.Now()
	balance := Balance{
		ID:        id.New(),
		UserID:    dto.UserID,
		Basic:     dto.Basic,
		Fun:       dto.Fun,
		Saves:     dto.Saves,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// TODO: validate numbers

	err := s.store.Insert(ctx, balance)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (s *balanceService) UpdateBalance(ctx context.Context, dto BalanceUpdate) error {
	if err := validate.Check(dto); err != nil {
		return err
	}

	// TODO: validate numbers

	err := s.store.Update(ctx, dto)
	return err
}

func (s *balanceService) GetBalance(ctx context.Context, userID string) (*Balance, error) {
	b, err := s.store.Get(ctx, userID)
	if err != nil {
		return nil, errors.NotFound("balance not found")
	}

	return b, nil
}
