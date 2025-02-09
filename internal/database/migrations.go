package database

import (
	"github.com/henriquepw/pobrin-api/internal/domains/user"
	"github.com/jmoiron/sqlx"
)

type Migration func(*sqlx.DB) error

func UserMigration(db *sqlx.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
      id TEXT PRIMARY KEY,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL,
	  deleted_at DATETIME NULL,
	  last_login DATETIME NULL,
      name TEXT NOT NULL,
      username TEXT NOT NULL,
      email TEXT NOT NULL,
      password TEXT NOT NULL
    )
  `

	_, err := db.Exec(schema)

	return err
}

func MainMigration(db *sqlx.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS income (
      id TEXT PRIMARY KEY,
      amount INTEGER NOT NULL,
      received_at DATETIME NOT NULL,
      created_at DATETIME NOT NULL,
      updated_at DATETIME NOT NULL
    )
  `

	_, err := db.Exec(schema)

	return err
}

func InsertUser(createdUser *user.User) func(db *sqlx.DB) error {
	return func(db *sqlx.DB) error {
		_, err := db.NamedExec(`
	insert into users 
	(id,  name,  username,  email,  password,  created_at,  updated_at, deleted_at)
		values
	(:id, :name, :username, :email, :password, :created_at, :updated_at, :deleted_at)
	`, createdUser)

		return err
	}
}
