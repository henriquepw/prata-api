// Package balance handle how the user income is divided, a goal to be achieved at time
package balance

import (
	"time"

	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/id"
)

type PieceUpdate struct {
	ID      id.ID  `json:"id" validate:"omitempty"`
	Label   string `json:"label" validate:"required,min=3"`
	Percent int    `json:"percent" validate:"min=0,max=100"`
	Color   string `json:"color" validate:"required"`
}

type BalanceUpdate struct {
	UserID id.ID         `json:"userId" validate:"required"`
	Pieces []PieceUpdate `json:"pieces" validate:"required"`
}

type Piece struct {
	ID        id.ID      `json:"id" db:"id"`
	UserID    id.ID      `json:"userId" db:"user_id"`
	Label     string     `json:"label" db:"label"`
	Color     string     `json:"color" db:"color"`
	Percent   int        `json:"percent" db:"percent"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type Balance struct {
	Pieces []Piece `json:"pieces"`
}

func (b Balance) CheckPercent() error {
	var percentTotal int
	for _, p := range b.Pieces {
		percentTotal += p.Percent
	}

	if percentTotal != 100 {
		return errorx.InvalidRequestData(map[string]string{
			"percent": "the sum of percent pieces should be 100",
		})
	}

	return nil
}
