package recurrence

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
)

type Recurrence struct {
	ID           id.ID      `json:"id" db:"id"`
	AccountID    id.ID      `json:"accountId" db:"account_id"`
	Description  string     `json:"description" db:"description"`
	Frequence    Frequence  `json:"frequence" db:"frequence"`
	Installments uint       `json:"installments" db:"installments"`
	StartAt      time.Time  `json:"startAt" db:"start_at"`
	EndAt        *time.Time `json:"endAt,omitempty" db:"end_at"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}

type RecurrenceCreate struct {
	AccountID    id.ID      `json:"accountId" validate:"required"`
	Description  string     `json:"description" validate:"required"`
	Frequence    Frequence  `json:"frequence" validate:"required,custom"`
	Installments uint       `json:"installments" validate:"required,min=1"`
	StartAt      time.Time  `json:"startAt" validate:"required"`
	EndAt        *time.Time `json:"endAt" validate:"omitempty,required"`
}

type RecurrenceUpdate struct {
	Description  *string    `json:"description" validate:"omitempty"`
	Frequence    *Frequence `json:"frequence" validate:"omitempty,custom"`
	Installments *uint      `json:"installments" validate:"omitempty,min=1"`
	EndAt        *time.Time `json:"endAt" validate:"omitempty"`
}

type RecurrenceQuery struct {
	Cursor    string
	Limit     int
	Frequence string
	Search    string
}
