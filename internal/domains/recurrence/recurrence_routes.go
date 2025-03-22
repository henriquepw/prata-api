package recurrence

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	store := NewStore(db)
	svc := NewService(store)
	handler := NewHandler(svc)

	return func(r chi.Router) {
		r.Post("/", handler.PostRecurrence)
		r.Get("/", handler.GetRecurrenceList)
		r.Get("/{id}", handler.GetRecurrenceByID)
		r.Patch("/{id}", handler.PatchRecurrenceByID)
		r.Delete("/{id}", handler.DeleteRecurrenceByID)
	}
}
