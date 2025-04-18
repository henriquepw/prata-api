package id

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/henriquepw/pobrin-api/pkg/errorx"
	cuid "github.com/nrednav/cuid2"
)

var ErrInvalidID = errorx.ServerError{Message: "invalid format id"}

type ID string

func isValid[T string | ID](id T) bool {
	return cuid.IsCuid(string(id))
}

func Parse(s string) (ID, error) {
	if isValid(s) {
		return "", ErrInvalidID
	}

	return ID(s), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if isValid(s) {
		return ErrInvalidID
	}

	*id = ID(s)
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	if isValid(string(id)) {
		return nil, ErrInvalidID
	}

	return json.Marshal(string(id))
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

func mustCreateID(size int) ID {
	generate, err := cuid.Init(
		cuid.WithLength(size),
		cuid.WithFingerprint("pobrin-api"),
	)
	if err != nil {
		panic(err)
	}

	return ID(generate())
}

func New() ID {
	return mustCreateID(24)
}

func NewTiny() ID {
	return mustCreateID(6)
}
