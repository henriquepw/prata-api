package balance

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type BalanceService interface {
	UpsertBalance(ctx context.Context, dto BalanceUpdate) (*Balance, error)
	GetBalance(ctx context.Context, userID string) (*Balance, error)
}

type balanceService struct {
	store BalanceStore
}

func NewService(store BalanceStore) BalanceService {
	return &balanceService{store}
}

func (s *balanceService) UpsertBalance(ctx context.Context, dto BalanceUpdate) (*Balance, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	balance, err := s.GetBalance(ctx, dto.UserID)
	if err != nil {
		return nil, err
	}

	// TODO: validate existis data

	now := time.Now()
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
		return nil, err
	}

	err = s.store.Upsert(ctx, *balance)
	return balance, err
}

func (s *balanceService) GetBalance(ctx context.Context, userID string) (*Balance, error) {
	balance, err := s.store.Get(ctx, userID)
	if err != nil {
		return nil, errorx.NotFound("balance not found")
	}

	return &balance, nil
}
