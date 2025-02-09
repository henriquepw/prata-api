package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/henriquepw/pobrin-api/internal/database"
	"github.com/henriquepw/pobrin-api/internal/domains/auth"
	"github.com/henriquepw/pobrin-api/internal/domains/user"
	"github.com/henriquepw/pobrin-api/pkg/hash"
	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/testutil"
	"github.com/henriquepw/pobrin-api/pkg/testutil/assert"
)

func TestService(t *testing.T) {
	ctx := context.Background()

	createdUser := &user.User{
		ID:       id.New(),
		Name:     "Nicole",
		Username: "nicole.chupetinha",
		Email:    "nicole.chupetinha@mail.com",
		Password: hash.MustGenerate("senha"),
	}

	t.Run("Success case", func(t *testing.T) {
		db := testutil.GetDB(database.UserMigration, database.InsertUser(createdUser))
		store := user.NewUserStore(db)
		service := auth.NewService(store)

		data := &auth.LoginRequest{
			Username: createdUser.Username,
			Password: "senha",
		}

		response, err := service.Login(ctx, data)
		assert.Nil(t, err)
		assert.NotEmptyString(t, response.Refresh)
		assert.NotEmptyString(t, response.Access)
	})

	t.Run("Not found user", func(t *testing.T) {
		db := testutil.GetDB(database.UserMigration)
		store := user.NewUserStore(db)
		service := auth.NewService(store)

		data := &auth.LoginRequest{
			Username: createdUser.Username,
			Password: "senha",
		}

		_, err := service.Login(ctx, data)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &user.ErrUserNotFound))
	})

	t.Run("Invalid Password", func(t *testing.T) {
		db := testutil.GetDB(database.UserMigration, database.InsertUser(createdUser))
		store := user.NewUserStore(db)
		service := auth.NewService(store)

		data := &auth.LoginRequest{
			Username: createdUser.Username,
			Password: "senhaErrada",
		}

		_, err := service.Login(ctx, data)
		assert.NotNil(t, err)
		assert.True(t, errors.As(err, &auth.ErrInvalidPassword))
	})
}
