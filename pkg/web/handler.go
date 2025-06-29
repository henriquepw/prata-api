package web

import (
	"net/http"
)

type HandlerFn func(w http.ResponseWriter, r *http.Request) error

func mainHandler(h HandlerFn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			ErrorResponse(w, err)
		}
	}
}
