package recurrence

import "github.com/go-chi/chi/v5"

func NewRouter() func(r chi.Router) {
	return func(r chi.Router) {}
}
