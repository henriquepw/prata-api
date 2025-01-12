package httputil

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/henriquepw/pobrin-api/pkg/errors"
)

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

func SuccessResponse(w http.ResponseWriter, data ...any) {
	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	writeJSON(w, http.StatusOK, data[0])
}

type Created struct {
	ID string `json:"id"`
}

func SuccessCreatedResponse(w http.ResponseWriter, id string) {
	writeJSON(w, http.StatusCreated, &Created{id})
}

func ErrorResponse(w http.ResponseWriter, err error) {
	slog.Error("HTTP API", "err", err.Error())

	if e, ok := err.(errors.ServerError); ok {
		writeJSON(w, e.StatusCode, e)
		return
	}

	internalErr := errors.Internal()
	writeJSON(w, internalErr.StatusCode, internalErr)
}

func CustomErrorResponse(w http.ResponseWriter, statusCode int, data any) {
	writeJSON(w, statusCode, data)
}
