package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquepw/pobrin-api/internal/domains/auth/session"
	"github.com/henriquepw/pobrin-api/internal/domains/auth/user"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	userStore := user.NewStore(db)
	userSVC := user.NewService(userStore)

	sessionStore := session.NewStore(db)
	sessionSVC := session.NewService(sessionStore)

	svc := NewService(userSVC, sessionSVC)
	handler := NewHandler(svc)

	return func(r chi.Router) {
		r.Post("/sign-in", handler.PostSignIn)
		r.Post("/sign-up", handler.PostSignUp)
		r.Post("/rewew/:token", handler.PostRenew)
	}
}
