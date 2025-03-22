package recurrence

import "time"

type Recurrence struct {
	ID string `json:"id"`
}

type RecurrenceCreate struct {
	Amount     int       `json:"amount" validate:"required,min=0"`
	ReceivedAt time.Time `json:"receivedAt" validate:"required"`
}

type RecurrenceUpdate struct {
	Amount     int       `json:"amount"`
	ReceivedAt time.Time `json:"receivedAt"`
}

type RecurrenceQuery struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit" validate:"required,min=0"`
}
