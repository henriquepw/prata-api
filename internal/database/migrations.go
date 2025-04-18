package database

import (
	"github.com/jmoiron/sqlx"
)

type Migration func(*sqlx.DB) error

func BalanceMigration(db *sqlx.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS balances (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      label INTEGER NOT NULL,
      percent DATETIME NOT NULL,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL
    );
  `

	_, err := db.Exec(schema)
	return err
}

func RecurrenceMigration(db *sqlx.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS recurrences (
      id TEXT PRIMARY KEY,
      description TEXT NOT NULL,
      frequence TEXT NOT NULL,
      installments INTEGER NOT NULL,
      start_at DATETIME NOT NULL,
      end_at DATETIME,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL,
      deleted_at DATETIME
    );
  `

	_, err := db.Exec(schema)
	return err
}

func TransactionMigration(db *sqlx.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS transactions (
      id TEXT PRIMARY KEY,
      amount INTEGER NOT NULL,
      received_at DATETIME NOT NULL,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL
    );
  `

	_, err := db.Exec(schema)
	return err
}
