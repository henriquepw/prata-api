package user

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/hash"
	"github.com/henriquepw/prata-api/pkg/id"
)

type UserService interface {
	CreateUser(ctx context.Context, email, password string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type userService struct {
	store UserStore
}

func NewService(store UserStore) UserService {
	return &userService{store}
}

func (s *userService) CreateUser(ctx context.Context, email, password string) (*User, error) {
	// check if user email already as registrated
	exists, err := s.store.Has(ctx, email)
	if err != nil {
		log.Error(err)
		return nil, errorx.Internal()
	}
	if exists {
		return nil, errorx.Conflict("Email already registrated")
	}

	// hash the user password
	secret, err := hash.Generate(password)
	if err != nil {
		return nil, errorx.Internal()
	}

	// create user into db
	user := User{
		ID:     id.New(),
		Email:  email,
		Secret: secret,
	}

	err = s.store.Insert(ctx, user)
	if err != nil {
		return nil, errorx.Internal("Can't create user")
	}

	return &user, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.store.Get(ctx, email)
	if err != nil {
		return nil, errorx.NotFound("user not found")
	}

	return user, nil
}
