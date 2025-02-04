package user_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/henriquepw/pobrin-api/internal/database"
	"github.com/henriquepw/pobrin-api/internal/domains/user"
	"github.com/henriquepw/pobrin-api/pkg/hash"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/testutil"
	"github.com/henriquepw/pobrin-api/pkg/testutil/assert"
	"github.com/jmoiron/sqlx"
)

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

func TestStore(t *testing.T) {
	ctx := context.Background()
	t.Run("Create user", func(t *testing.T) {
		db := testutil.GetDB(database.UserMigration)
		userStore := user.NewUserStore(db)

		createdUser := &user.User{ID: id.New(), Name: "Jamiro", Email: "jamiro@mail.com", Username: "XjamiroX", Password: hash.MustGenerate("secure")}

		err := userStore.Insert(ctx, createdUser)
		assert.Nil(t, err)

		foundUser := user.User{}
		err = db.Get(&foundUser, "select * from users limit 1")
		t.Log(foundUser)
		assert.Nil(t, err)
		assert.Equal(t, createdUser.ID, foundUser.ID)
	})

	t.Run("Find user", func(t *testing.T) {
		createdUser := &user.User{ID: id.New(), Name: "Jamiro", Email: "jamiro@mail.com", Username: "XjamiroX", Password: hash.MustGenerate("secure")}

		db := testutil.GetDB(database.UserMigration, InsertUser(createdUser))
		userStore := user.NewUserStore(db)

		foundUser, err := userStore.Get(ctx, createdUser.ID)
		assert.Nil(t, err)
		assert.Equal(t, createdUser.ID, foundUser.ID)
		assert.Equal(t, createdUser.Email, foundUser.Email)
	})

	t.Run("Get user password", func(t *testing.T) {
		createdUser := &user.User{ID: id.New(), Name: "Jamiro", Email: "jamiro@mail.com", Username: "XjamiroX", Password: hash.MustGenerate("secure")}

		db := testutil.GetDB(database.UserMigration, InsertUser(createdUser))
		userStore := user.NewUserStore(db)

		password, err := userStore.GetUserPassword(ctx, createdUser.Username)
		assert.Nil(t, err)
		assert.True(t, hash.Validate(password, "secure"))
	})

	t.Run("Get user password", func(t *testing.T) {
		db := testutil.GetDB(database.UserMigration)
		userStore := user.NewUserStore(db)

		_, err := userStore.GetUserPassword(ctx, "notfound")
		assert.NotNil(t, err)
	})

	t.Run("Find deleted user", func(t *testing.T) {
		createdUser := &user.User{ID: id.New(), Name: "Jamiro", Email: "jamiro@mail.com", Username: "XjamiroX", Password: hash.MustGenerate("secure"), DeletedAt: sql.NullTime{Time: time.Now(), Valid: true}}

		db := testutil.GetDB(database.UserMigration, InsertUser(createdUser))
		userStore := user.NewUserStore(db)

		_, err := userStore.Get(ctx, createdUser.ID)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &user.UserNotFound))
	})

	t.Run("Delete user", func(t *testing.T) {
		createdUser := &user.User{ID: id.New(), Name: "Jamiro", Email: "jamiro@mail.com", Username: "XjamiroX", Password: hash.MustGenerate("secure")}

		db := testutil.GetDB(database.UserMigration, InsertUser(createdUser))
		userStore := user.NewUserStore(db)

		err := userStore.Delete(ctx, createdUser.ID)
		assert.Nil(t, err)
		foundUser := user.User{}
		err = db.Get(&foundUser, "select * from users limit 1")
		assert.Nil(t, err)
		assert.True(t, foundUser.DeletedAt.Valid)
		assert.TimeIsNotZero(t, foundUser.DeletedAt.Time)
	})

	t.Run("Update user", func(t *testing.T) {
		createdUser := &user.User{ID: id.New(), Name: "Jamiro", Email: "jamiro@mail.com", Username: "XjamiroX", Password: hash.MustGenerate("secure")}

		db := testutil.GetDB(database.UserMigration, InsertUser(createdUser))
		userStore := user.NewUserStore(db)

		createdUser.Username = "jamiro.gameplays"

		err := userStore.Update(ctx, createdUser.ID, createdUser)
		assert.Nil(t, err)

		foundUser := &user.User{}
		db.Get(foundUser, "select * from users")
		assert.Equal(t, createdUser.Username, foundUser.Username)
	})
}
