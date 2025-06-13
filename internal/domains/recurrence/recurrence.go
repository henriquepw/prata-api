package recurrence

import (
	"time"

	"github.com/henriquepw/prata-api/internal/domains/transaction"
	"github.com/henriquepw/prata-api/pkg/id"
)

type Recurrence struct {
	ID          id.ID                       `json:"id" db:"id"`
	UserID      id.ID                       `json:"userId" db:"user_id"`
	BalanceID   *id.ID                      `json:"balanceId" db:"balance_id"`
	Amount      int                         `json:"amount" db:"amount"`
	Description string                      `json:"description" db:"description"`
	Type        transaction.TransactionType `json:"type" db:"type"`
	Frequence   Frequence                   `json:"frequence" db:"frequence"`
	StartAt     time.Time                   `json:"startAt" db:"start_at"`
	EndAt       *time.Time                  `json:"endAt,omitempty" db:"end_at"`
	Day         int                         `json:"day" db:"day"`
	Week        int                         `json:"week" db:"week"`
	Month       int                         `json:"month" db:"month"`
	YearDay     int                         `json:"yearDay" db:"year_day"`
	CreatedAt   time.Time                   `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time                   `json:"updatedAt" db:"updated_at"`
}

type RecurrenceCreate struct {
	UserID      id.ID                       `json:"userID" validate:"required"`
	BalanceID   *id.ID                      `json:"balanceId"`
	Amount      int                         `json:"amount" validate:"required"`
	Description string                      `json:"description" validate:"required"`
	Frequence   Frequence                   `json:"frequence" validate:"required,custom"`
	Type        transaction.TransactionType `json:"type" validate:"required,custom"`
	StartAt     time.Time                   `json:"startAt" validate:"required"`
	EndAt       *time.Time                  `json:"endAt" validate:"omitempty,required"`
}

type RecurrenceUpdate struct {
	BalanceID   *id.ID     `json:"balanceId" validate:"omitempty"`
	Amount      int        `json:"amount" validate:"omitempty"`
	Description string     `json:"description" validate:"omitempty"`
	Frequence   Frequence  `json:"frequence" validate:"omitempty,custom"`
	EndAt       *time.Time `json:"endAt" validate:"omitempty"`
}

type RecurrenceQuery struct {
	Cursor     string
	Limit      int
	UserID     id.ID
	Search     string
	Frequence  Frequence
	Type       transaction.TransactionType
	EndAtGte   time.Time
	EndAtLte   time.Time
	StartAtGte time.Time
	StartAtLte time.Time
}
