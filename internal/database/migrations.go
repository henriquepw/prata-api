package database

import (
	"github.com/jmoiron/sqlx"
)

type Migration func(*sqlx.DB) error

func BalanceMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS balances (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      label INTEGER NOT NULL,
      percent DATETIME NOT NULL,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL
    );
  `)

	db.MustExec(`
		CREATE INDEX idx_balances_user_id
		ON balances (user_id);
	`)

	return nil
}

func RecurrenceMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS recurrences (
      id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			balance_id TEXT,
	    amount INTEGER NOT NULL,
      description TEXT NOT NULL,
      frequence TEXT NOT NULL,
      type TEXT NOT NULL,
      day INTEGER NOT NULL,
      week INTEGER NOT NULL,
      month INTEGER NOT NULL,
      year_day INTEGER NOT NULL,
      start_at DATETIME NOT NULL,
      end_at DATETIME,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL,
			FOREIGN KEY(balance_id) REFERENCES balances(id)
    );
  `)

	db.MustExec(`
		CREATE INDEX idx_recurrences_frequence_end_at
		ON recurrences (frequence, end_at);

		CREATE INDEX idx_recurrences_week
		ON recurrences (week);

		CREATE INDEX idx_recurrences_year_day
		ON recurrences (year_day);

		CREATE INDEX idx_recurrences_day_month
		ON recurrences (day, month);

		CREATE INDEX idx_recurrences_user_id
		ON recurrences (user_id);
	`)

	return nil
}

func TransactionMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS transactions (
      id TEXT PRIMARY KEY,
      user_Id TEXT NOT NULL,
      balance_id TEXT,
      type TEXT NOT NULL,
      description TEXT NOT NULL,
      amount INTEGER NOT NULL,
      received_at DATETIME NOT NULL,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL,
			FOREIGN KEY(balance_id) REFERENCES balances(id)
    );
  `)

	db.MustExec(`
		CREATE INDEX idx_transactions_user_id
		ON transactions (user_id);
	`)

	return nil
}
