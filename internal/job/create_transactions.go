package job

import (
	"context"
	"time"

	"github.com/henriquepw/prata-api/internal/domains/transaction"
)

func (s *jobServer) createTransactionByRecurrence() error {
	ctx := context.TODO()

	recurrence := s.recurrenceStore.TodayRecurrences(ctx)
	if len(recurrence) == 0 {
		return nil
	}

	now := time.Now()
	payload := []transaction.TransactionCreate{}
	for _, t := range recurrence {
		payload = append(payload, transaction.TransactionCreate{
			UserID:      t.UserID,
			BalanceID:   t.BalanceID,
			Amount:      t.Amount,
			Type:        t.Type,
			Description: t.Description,
			ReceivedAt:  now,
		})
	}

	_, err := s.transactionSVC.CreateManyTransactions(ctx, payload)
	return err
}
