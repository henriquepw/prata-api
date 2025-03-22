package transaction

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
)

const (
	TypeIncome  = "INCOME"
	TypeOutcome = "OUTCOME"
)

type TransactionType string

func (f TransactionType) Validate() bool {
	switch f {
	case TypeIncome, TypeOutcome:
		return true
	}
	return false
}

type Transaction struct {
	ID          id.ID           `json:"id" db:"id"`
	AccountID   id.ID           `json:"accountId" db:"account_id"`
	Tags        []string        `json:"tags" db:"tags"`
	Type        TransactionType `json:"type" db:"type"`
	Description string          `json:"description" db:"description"`
	Amount      int             `json:"amount" db:"amount"`
	ReceivedAt  time.Time       `json:"receivedAt" db:"received_at"`
	CreatedAt   time.Time       `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time       `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time      `json:"deletedAt" db:"deleted_at"`
}

type TransactionCreate struct {
	AccountID   id.ID           `json:"accountId" validate:"required"`
	Tags        []string        `json:"tags" validate:"required"`
	Type        TransactionType `json:"type" validate:"required,custom"`
	Description string          `json:"description" validate:"required"`
	Amount      int             `json:"amount" validate:"required"`
	ReceivedAt  time.Time       `json:"receivedAt" validate:"required"`
}

type TransactionUpdate struct {
	Amount     int       `json:"amount"`
	ReceivedAt time.Time `json:"receivedAt"`
}

type TransactionQuery struct {
	Cursor    string `json:"cursor"`
	Limit     int    `json:"limit" validate:"required,min=0"`
	AccountID string `json:"accountId"`
}
