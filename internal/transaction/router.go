package transaction

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	store := NewStore(db)
	svc := NewService(store)
	handler := NewHandler(svc)

	return func(r chi.Router) {
		r.Post("/", handler.PostTransaction)
		r.Get("/", handler.GetTransactionList)
		r.Get("/{id}", handler.GetTransactionByID)
		r.Patch("/{id}", handler.PatchTransactionByID)
		r.Delete("/{id}", handler.DeleteTransactionByID)
	}
}
