package web

import (
	"encoding/json"
	"net/http"

	"github.com/henriquepw/prata-api/pkg/errorx"
)

func GetBodyRequest[T any](r *http.Request) (T, error) {
	defer r.Body.Close()

	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, errorx.InvalidJSON()
	}

	return data, nil
}
