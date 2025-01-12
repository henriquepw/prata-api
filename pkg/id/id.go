package id

import (
	"database/sql/driver"
	"encoding/json"

	serverError "github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/nrednav/cuid2"
)

var (
	ErrInvalidID = serverError.ServerError{Message: "invalid format id"}
)

type ID string

func New() ID {
	return createID(24)
}

func NewTiny() ID {
	return createID(6)
}

func createID(size int) ID {
	generate, err := cuid2.Init(
		cuid2.WithLength(size),
		cuid2.WithFingerprint("pobrin-api"),
	)
	if err != nil {
		panic(err)
	}

	return ID(generate())
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	ok := cuid2.IsCuid(s)
	if !ok {
		return ErrInvalidID
	}

	*id = ID(s)

	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	var s string

	ok := cuid2.IsCuid(string(id))
	if !ok {
		return nil, ErrInvalidID
	}

	s = string(id)

	return json.Marshal(s)
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
