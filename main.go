package main

import (
	"log/slog"

	"github.com/henriquepw/prata-api/internal/api"
	"github.com/henriquepw/prata-api/internal/database"
	"github.com/henriquepw/prata-api/internal/job"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := database.GetDB()
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		return
	}
	defer db.Close()

	jobServer := job.New(db)
	if err := jobServer.Start(); err != nil {
		slog.Error("failed to start job server", "error", err)
		return
	}

	apiServer := api.New(db)
	if err := apiServer.Start(); err != nil {
		slog.Error("failed to initialize api server", "error", err)
		return
	}
}
