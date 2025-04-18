package balance

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type BalanceService interface {
	CreateBalance(ctx context.Context, dto BalanceCreate) (Balance, error)
	UpdateBalance(ctx context.Context, dto BalanceUpdate) (Balance, error)
	GetBalance(ctx context.Context, userID string) (Balance, error)
}

type balanceService struct {
	store BalanceStore
}

func NewService(store BalanceStore) BalanceService {
	return &balanceService{store}
}

func (s *balanceService) CreateBalance(ctx context.Context, dto BalanceCreate) (Balance, error) {
	now := time.Now()
	balance := Balance{}

	if err := validate.Check(dto); err != nil {
		return balance, err
	}

	// TODO: validate existis data

	for _, p := range dto.Pieces {
		balance.Pieces = append(balance.Pieces, Piece{
			ID:        id.New(),
			UserID:    dto.UserID,
			Label:     p.Label,
			Percent:   p.Percent,
			CreatedAt: now,
			UpdatedAt: now,
		})
	}

	if err := balance.CheckPercent(); err != nil {
		return balance, err
	}

	err := s.store.Upsert(ctx, balance)
	return balance, err
}

func (s *balanceService) UpdateBalance(ctx context.Context, dto BalanceUpdate) (Balance, error) {
	now := time.Now()
	balance := Balance{}

	if err := validate.Check(dto); err != nil {
		return balance, err
	}

	// TODO: validate existis data

	for _, p := range dto.Pieces {
		piece := Piece{
			ID:        p.ID,
			UserID:    dto.UserID,
			Label:     p.Label,
			Percent:   p.Percent,
			CreatedAt: now,
			UpdatedAt: now,
		}
		if piece.ID == "" {
			piece.ID = id.New()
		}

		balance.Pieces = append(balance.Pieces, piece)
	}

	if err := balance.CheckPercent(); err != nil {
		return balance, err
	}

	err := s.store.Upsert(ctx, balance)
	return balance, err
}

func (s *balanceService) GetBalance(ctx context.Context, userID string) (Balance, error) {
	balance, err := s.store.Get(ctx, userID)
	if err != nil {
		return balance, errorx.NotFound("balance not found")
	}

	return balance, nil
}
