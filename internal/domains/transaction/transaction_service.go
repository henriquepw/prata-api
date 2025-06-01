package transaction

import (
	"context"
	"time"

	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/id"
	"github.com/henriquepw/prata-api/pkg/page"
	"github.com/henriquepw/prata-api/pkg/validate"
)

type TransactionService interface {
	CreateTransactions(ctx context.Context, dto []TransactionCreate) ([]Transaction, error)
	UpdateTransaction(ctx context.Context, id id.ID, dto TransactionUpdate) error
	DeleteTransaction(ctx context.Context, id id.ID) error
	GetTransaction(ctx context.Context, id id.ID) (Transaction, error)
	ListTransaction(ctx context.Context, dto TransactionQuery) page.Cursor[Transaction]
}

type transactionService struct {
	store TransactionStore
}

func NewService(store TransactionStore) TransactionService {
	return &transactionService{store}
}

func (s *transactionService) CreateTransactions(ctx context.Context, dto []TransactionCreate) ([]Transaction, error) {
	trxs := []Transaction{}
	if err := validate.Check(dto); err != nil {
		return trxs, err
	}

	now := time.Now()
	for _, item := range dto {
		trxs = append(trxs, Transaction{
			ID:          id.New(),
			UserID:      item.UserID,
			Description: item.Description,
			Type:        item.Type,
			Amount:      item.Amount,
			ReceivedAt:  item.ReceivedAt,
			CreatedAt:   now,
			UpdatedAt:   now,
		})
	}

	err := s.store.Insert(ctx, trxs)
	if err != nil {
		return trxs, errorx.Internal("Failed to create the transaction")
	}

	return trxs, nil
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

func (s *transactionService) GetTransaction(ctx context.Context, id id.ID) (Transaction, error) {
	transaction, err := s.store.Get(ctx, id)
	if err != nil {
		return transaction, errorx.NotFound("Transaction not found")
	}

	return transaction, nil
}

func (s *transactionService) ListTransaction(ctx context.Context, dto TransactionQuery) page.Cursor[Transaction] {
	if err := validate.Check(dto); err != nil {
		return page.NewEmpty[Transaction]()
	}

	items, err := s.store.List(ctx, dto)
	if err != nil {
		return page.NewEmpty[Transaction]()
	}

	return items
}
