package id

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/oklog/ulid/v2"
)

var ErrInvalidID = errorx.BadRequest("invalid format id")

type ID string

func isValid[T string | ID](id T) bool {
	_, err := ulid.Parse(string(id))
	return err == nil
}

func Parse(s string) (ID, error) {
	if !isValid(s) {
		return "", ErrInvalidID
	}

	return ID(s), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*id = ID(s)
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	if !isValid(id) {
		return nil, ErrInvalidID
	}

	return json.Marshal(id.String())
}

func (id ID) String() string {
	return string(id)
}

func (id *ID) Scan(value any) error {
	if value == nil {
		*id = ""
		return nil
	}

	switch v := value.(type) {
	case string:
		*id = ID(v)
	case []byte:
		*id = ID(string(v))
	case ID:
		*id = v
	default:
		return ErrInvalidID
	}

	return nil
}

func (id ID) Value() (driver.Value, error) {
	if id == "" {
		return nil, nil
	}

	return string(id), nil
}

func New() ID {
	return ID(ulid.Make().String())
}
