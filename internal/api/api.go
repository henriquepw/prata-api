package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/henriquepw/pobrin-api/internal/config"
	"github.com/henriquepw/pobrin-api/internal/domains/income"
	"github.com/henriquepw/pobrin-api/pkg/errors"
	"github.com/henriquepw/pobrin-api/pkg/httputil"
	"github.com/jmoiron/sqlx"
)

type apiServer struct {
	db   *sqlx.DB
	addr string
}

func NewApiServer(db *sqlx.DB) *apiServer {
	return &apiServer{db, ":" + config.Env().Port}
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
		middleware.AllowContentType("application/json", "application/x-www-form-urlencoded", "multipart/form-data"),
		middleware.Heartbeat("/health"),
	)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		httputil.ErrorResponse(w, errors.NotFound("rota não encontrada"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		httputil.ErrorResponse(w, errors.MethodNotAllowed())
	})

	incomeStore := income.NewIncomeStore(s.db)
	incomeSvc := income.NewIncomeService(incomeStore)
	incomeHandler := income.NewIncomeHandler(incomeSvc)

	// Private Routes
	r.Group(func(r chi.Router) {
		// add auth middleware

		r.Route("/incomes", func(r chi.Router) {
			r.Post("/", incomeHandler.PostIncome)
			r.Get("/", incomeHandler.GetIncomeList)
			r.Get("/{incomeId}", incomeHandler.GetIncomeByID)
			r.Patch("/{incomeId}", incomeHandler.PatchIncomeByID)
			r.Delete("/{incomeId}", incomeHandler.DeleteIncomeByID)
		})
	})

	return http.ListenAndServe(s.addr, r)
}
