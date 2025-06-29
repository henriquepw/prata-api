package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/henriquepw/prata-api/internal/database"
	"github.com/henriquepw/prata-api/internal/domains/auth"
	"github.com/henriquepw/prata-api/internal/domains/balance"
	"github.com/henriquepw/prata-api/internal/domains/recurrence"
	"github.com/henriquepw/prata-api/internal/domains/transaction"
	"github.com/henriquepw/prata-api/internal/env"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/httpx"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type apiServer struct {
	ctx  context.Context
	addr string
	db   *sqlx.DB
}

func New(ctx context.Context, db *sqlx.DB) *apiServer {
	return &apiServer{ctx, ":" + os.Getenv(env.Port), db}
}

func (s *apiServer) Start() error {
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

	r.Group(auth.NewRouter(s.db))

	r.Route("/me", func(r chi.Router) {
		r.Use(auth.RequireAuthorization)
		r.Route("/balance", balance.NewRouter(s.db))
		r.Route("/transactions", transaction.NewRouter(s.db))
		r.Route("/recurrences", recurrence.NewRouter(s.db))
	})

	fmt.Println("Server running on port ", s.addr)
	return http.ListenAndServe(s.addr, r)
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	db, err := database.GetDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if os.Getenv(env.Version) == "DEVELOP" {
		db.SetMaxOpenConns(1)
	}

	apiServer := New(ctx, db)
	return apiServer.Start()
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error("Cronjob finished with error", "error", err.Error())
		os.Exit(1)
	}
}
