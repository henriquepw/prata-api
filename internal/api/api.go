package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/henriquepw/pobrin-api/internal/domains/balance"
	"github.com/henriquepw/pobrin-api/internal/domains/recurrence"
	"github.com/henriquepw/pobrin-api/internal/domains/transaction"
	"github.com/henriquepw/pobrin-api/internal/env"
	"github.com/henriquepw/pobrin-api/pkg/errorx"
	"github.com/henriquepw/pobrin-api/pkg/httpx"
	"github.com/jmoiron/sqlx"
)

type apiServer struct {
	db   *sqlx.DB
	addr string
}

func New(db *sqlx.DB) *apiServer {
	return &apiServer{db, ":" + os.Getenv(env.Port)}
}

func (s *apiServer) Start() error {
	clerk.SetKey(os.Getenv(env.ClerkApiKey))

	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
		middleware.AllowContentType("application/json"),
		middleware.Heartbeat("/health"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		httpx.ErrorResponse(w, errorx.NotFound("rota não encontrada"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		httpx.ErrorResponse(w, errorx.MethodNotAllowed())
	})

	r.Route("/user", func(r chi.Router) {
		r.Use(clerkhttp.WithHeaderAuthorization())

		r.Route("/balance", balance.NewRouter(s.db))
		r.Route("/transactions", transaction.NewRouter(s.db))
		r.Route("/recurrences", recurrence.NewRouter(s.db))
	})

	fmt.Println("Server running on port ", s.addr)
	return http.ListenAndServe(s.addr, r)
}
