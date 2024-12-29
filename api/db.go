package api

import (
	"log"

	"github.com/henriquepw/pobrin-api/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/go-libsql"
)

func StartDB() *sqlx.DB {
	db, err := sqlx.Open("libsql", config.Env().DatabaseURL)
	if err != nil {
		log.Fatalf("failed to open db: %s", err.Error())
	}

	schema := `
    CREATE TABLE IF NOT EXISTS income (
      id TEXT PRIMARY KEY,
      amount INTEGER NOT NULL,
      received_at DATETIME NOT NULL,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL
    )
  `

	db.MustExec(schema)

	return db
}
