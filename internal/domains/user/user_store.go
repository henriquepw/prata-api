package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	Insert(ctx context.Context, data *User) error
	Delete(ctx context.Context, id id.ID) error
	Update(ctx context.Context, id id.ID, data *User) error
	Get(ctx context.Context, id id.ID) (*User, error)
	// List(ctx context.Context, q User) (*page.Cursor[User], error) // ToDo: to add later
	GetUserPassword(ctx context.Context, username string) (string, error)
}

type store struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *store {
	return &store{db: db}
}

func (store *store) Insert(ctx context.Context, user *User) error {
	q := `
	insert into users 
	(id,  name,  username,  email,  password,  created_at,  updated_at)
	    values
	(:id, :name, :username, :email, :password, :created_at, :updated_at)
	`

	_, err := store.db.NamedExecContext(ctx, q, user)
	if err != nil {
		return fmt.Errorf("failed to crete user: %w", err)
	}

	return nil
}

func (store *store) Delete(ctx context.Context, id id.ID) error {
	_, err := store.db.ExecContext(ctx, "UPDATE users SET deleted_at = ? WHERE id = ?", time.Now(), id)
	return err
}

func (store *store) Update(ctx context.Context, id id.ID, user *User) error {
	q := `
    UPDATE users
    SET name = ?, username = ?, email = ?, password = ?, updated_at = ?
    WHERE id = ?`

	_, err := store.db.ExecContext(ctx, q,
		user.Name, user.Username, user.Email, user.Password, time.Now(),
		id,
	)

	return err
}

func (store *store) Get(ctx context.Context, id id.ID) (*User, error) {
	query := "SELECT * FROM users WHERE id = ? and deleted_at is null"

	var user User
	err := store.db.GetContext(ctx, &user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (store *store) GetUserPassword(ctx context.Context, username string) (string, error) {
	query := "SELECT password FROM users WHERE username = ? and deleted_at is null"

	var password string
	err := store.db.GetContext(ctx, &password, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return password, err
	}

	return password, nil
}
