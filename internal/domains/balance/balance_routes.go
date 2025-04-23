package balance

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquepw/pobrin-api/internal/auth"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	store := NewStore(db)
	svc := NewService(store)
	handler := NewHandler(svc, auth.NewSession())

	return func(r chi.Router) {
		r.Post("/", handler.PostUserBalance)
		r.Get("/", handler.GetUserBalance)
	}
}
