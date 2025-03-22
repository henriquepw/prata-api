package transaction

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, dto TransactionCreate) (*Transaction, error)
	UpdateTransaction(ctx context.Context, id id.ID, dto TransactionUpdate) error
	DeleteTransaction(ctx context.Context, id id.ID) error
	GetTransaction(ctx context.Context, id id.ID) (*Transaction, error)
	ListTransaction(ctx context.Context, dto TransactionQuery) *page.Cursor[Transaction]
}

type transactionService struct {
	store TransactionStore
}

func NewService(store TransactionStore) TransactionService {
	return &transactionService{store}
}

func (s *transactionService) CreateTransaction(ctx context.Context, dto TransactionCreate) (*Transaction, error) {
	if err := validate.Check(dto); err != nil {
		return nil, err
	}

	now := time.Now()
	transaction := Transaction{
		ID:          id.New(),
		AccountID:   dto.AccountID,
		Description: dto.Description,
		Tags:        dto.Tags,
		Type:        dto.Type,
		Amount:      dto.Amount,
		ReceivedAt:  dto.ReceivedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.store.Insert(ctx, transaction)
	if err != nil {
		return nil, errors.Internal("Failed to create the transaction")
	}

	return &transaction, nil
}

func (s *transactionService) UpdateTransaction(ctx context.Context, id id.ID, dto TransactionUpdate) error {
	if err := validate.Check(dto); err != nil {
		return err
	}

	return nil
}

func (s *transactionService) DeleteTransaction(ctx context.Context, id id.ID) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return errors.Internal("Failed to delete the transaction")
	}

	return nil
}

func (s *transactionService) GetTransaction(ctx context.Context, id id.ID) (*Transaction, error) {
	transaction, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errors.NotFound("Transaction not found")
	}

	return transaction, nil
}

func (s *transactionService) ListTransaction(ctx context.Context, dto TransactionQuery) *page.Cursor[Transaction] {
	if err := validate.Check(dto); err != nil {
		return page.NewEmpty[Transaction]()
	}

	items, err := s.store.List(ctx, dto)
	if err != nil {
		return page.NewEmpty[Transaction]()
	}

	return items
}
