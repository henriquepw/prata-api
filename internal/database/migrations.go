package database

import (
	"github.com/jmoiron/sqlx"
)

type Migration func(*sqlx.DB) error

func UserMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS users (
      id TEXT PRIMARY KEY,
      email TEXT NOT NULL,
      username TEXT NOT NULL,
      avatar TEXT,
      secret INTEGER NOT NULL,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      deleted_at DATETIME
    );
  `)

	db.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_users_email
		ON users (email);
	`)

	return nil
}

func SessionMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS sessions (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      refresh_token TEXT NOT NULL,
      expires_at DATETIME NOT NULL,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
    );
  `)

	return nil
}

func BalanceMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS balances (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      label TEXT NOT NULL,
      color TEXT NOT NULL,
      percent INTEGER NOT NULL,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id)
    );
  `)

	db.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_balances_user_id
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
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(balance_id) REFERENCES balances(id)
    );
  `)

	db.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_recurrences_frequence_end_at
		ON recurrences (frequence, end_at);

		CREATE INDEX IF NOT EXISTS idx_recurrences_week
		ON recurrences (week);

		CREATE INDEX IF NOT EXISTS idx_recurrences_year_day
		ON recurrences (year_day);

		CREATE INDEX IF NOT EXISTS idx_recurrences_day_month
		ON recurrences (day, month);

		CREATE INDEX IF NOT EXISTS idx_recurrences_user_id
		ON recurrences (user_id);
	`)

	return nil
}

func TransactionMigration(db *sqlx.DB) error {
	db.MustExec(`
    CREATE TABLE IF NOT EXISTS transactions (
      id TEXT PRIMARY KEY,
      user_id TEXT NOT NULL,
      balance_id TEXT,
      type TEXT NOT NULL,
      description TEXT NOT NULL,
      amount INTEGER NOT NULL,
      received_at DATETIME NOT NULL,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users(id),
			FOREIGN KEY(balance_id) REFERENCES balances(id)
    );
  `)

	db.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_transactions_user_id
		ON transactions (user_id);
	`)

	return nil
}
