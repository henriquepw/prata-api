package transaction

import (
	"context"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/page"
	"github.com/henriquepw/pobrin-api/pkg/validate"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, dto TransactionCreate) (*Transaction, error)
	CreateManyTransactions(ctx context.Context, dto []TransactionCreate) ([]Transaction, error)
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
		UserID:      dto.UserID,
		Description: dto.Description,
		Type:        dto.Type,
		Amount:      dto.Amount,
		ReceivedAt:  dto.ReceivedAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.store.Insert(ctx, transaction)
	if err != nil {
		return nil, errorx.Internal("Failed to create the transaction")
	}

	return &transaction, nil
}

// TODO:
func (s *transactionService) CreateManyTransactions(ctx context.Context, dto []TransactionCreate) ([]Transaction, error) {
	transactions := []Transaction{}

	for _, t := range dto {
		item, err := s.CreateTransaction(ctx, t)
		if err != nil {
			continue
		}

		transactions = append(transactions, *item)
	}

	return transactions, nil
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
		return errorx.Internal("Failed to delete the transaction")
	}

	return nil
}

func (s *transactionService) GetTransaction(ctx context.Context, id id.ID) (*Transaction, error) {
	transaction, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, errorx.NotFound("Transaction not found")
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
