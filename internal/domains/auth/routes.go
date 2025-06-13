package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquepw/prata-api/internal/domains/auth/session"
	"github.com/henriquepw/prata-api/internal/domains/auth/user"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) func(r chi.Router) {
	userStore := user.NewStore(db)
	userSVC := user.NewService(userStore)

	sessionStore := session.NewStore(db)
	sessionSVC := session.NewService(sessionStore)

	authSVC := NewService(userSVC, sessionSVC)
	handler := NewHandler(authSVC, userSVC)

	return func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/sign-in", handler.PostSignIn)
			r.Post("/sign-up", handler.PostSignUp)
			r.Get("/renew/{token}", handler.PostRenew)
		})

		r.Group(func(r chi.Router) {
			r.Use(RequireAuthorization)
			r.Get("/me/profile", handler.GetUserProfile)
		})
	}
}
