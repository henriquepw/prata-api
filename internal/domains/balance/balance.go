package balance

import (
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
)

type Balance struct {
	ID        id.ID     `json:"id" db:"id"`
	UserID    string    `json:"userId" db:"user_id"`
	Basic     uint8     `json:"basic" db:"balance_basic" validate:"min=0,max=100"`
	Fun       uint8     `json:"fun" db:"balance_fun" validate:"min=0,max=100"`
	Saves     uint8     `json:"saves" db:"balance_saves" validate:"min=0,max=100"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type BalanceCreate struct {
	UserID string `json:"userId" validate:"required"`
	Basic  uint8  `json:"basic" validate:"min=0,max=100"`
	Fun    uint8  `json:"fun" validate:"min=0,max=100"`
	Saves  uint8  `json:"saves" validate:"min=0,max=100"`
}

type BalanceUpdate = BalanceCreate
