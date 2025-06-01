package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/henriquepw/prata-api/pkg/errorx"
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

func SuccessCreatedResponse(w http.ResponseWriter, data ...any) {
	writeJSON(w, http.StatusCreated, data)
}

func ErrorResponse(w http.ResponseWriter, err error) {
	log.Error("HTTP API", "err", err.Error())

	if e, ok := err.(errorx.ServerError); ok {
		writeJSON(w, e.StatusCode, e)
		return
	}

	internalErr := errorx.Internal()
	writeJSON(w, internalErr.StatusCode, internalErr)
}

func CustomErrorResponse(w http.ResponseWriter, statusCode int, data any) {
	writeJSON(w, statusCode, data)
}
