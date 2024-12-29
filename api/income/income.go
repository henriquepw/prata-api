package income

import "time"

type Income struct {
	ID         string    `json:"id" db:"id"`
	Amount     int       `json:"amount" db:"amount"`
	ReceivedAt time.Time `json:"receivedAt" db:"received_at"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

type IncomeCreate struct {
	Amount     int       `json:"amount" validate:"required,min=0"`
	ReceivedAt time.Time `json:"receivedAt" validate:"required"`
}

type IncomeUpdate struct {
	Amount     int       `json:"amount"`
	ReceivedAt time.Time `json:"receivedAt"`
}

type IncomeQuery struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit" validate:"required,min=0"`
}
