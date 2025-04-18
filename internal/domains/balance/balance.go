package balance

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/id"
)

type PieceCreate struct {
	Label   string `json:"label" validate:"required,min=3"`
	Percent uint8  `json:"percent" validate:"min=0,max=100"`
}

type BalanceCreate struct {
	UserID string        `json:"userId" validate:"required"`
	Pieces []PieceCreate `json:"pieces" validate:"required"`
}

type PieceUpdate struct {
	ID      id.ID  `json:"id"`
	Label   string `json:"label" validate:"required,min=3"`
	Percent uint8  `json:"percent" validate:"min=0,max=100"`
}

type BalanceUpdate struct {
	UserID string        `json:"userId" validate:"required"`
	Pieces []PieceUpdate `json:"pieces" validate:"required"`
}

type Piece struct {
	ID        id.ID     `json:"id" db:"id"`
	UserID    string    `json:"userId" db:"user_id"`
	Label     string    `json:"label" db:"label"`
	Percent   uint8     `json:"percent" db:"percent"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Balance struct {
	Pieces []Piece `json:"pieces"`
}

func (b Balance) CheckPercent() error {
	var percentTotal uint8
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
