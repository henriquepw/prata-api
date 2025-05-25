package user

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/henriquepw/prata-api/pkg/errorx"
	"github.com/henriquepw/prata-api/pkg/hash"
	"github.com/henriquepw/prata-api/pkg/id"
)

type UserService interface {
	CreateUser(ctx context.Context, dto UserCreate) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id id.ID) (User, error)
}

type userService struct {
	store UserStore
}

func NewService(store UserStore) UserService {
	return &userService{store}
}

func (s *userService) CreateUser(ctx context.Context, dto UserCreate) (User, error) {
	// check if user email already as registrated
	exists, err := s.store.Has(ctx, dto.Email)
	if err != nil {
		log.Error(err)
		return User{}, errorx.Internal()
	}
	if exists {
		return User{}, errorx.Conflict("Email already registrated")
	}

	// hash the user password
	secret, err := hash.Generate(dto.Password)
	if err != nil {
		return User{}, errorx.Internal()
	}

	// create user into db
	user := User{
		ID:       id.New(),
		Email:    dto.Email,
		Avatar:   dto.Avatar,
		Username: dto.Username,
		Secret:   secret,
	}

	err = s.store.Insert(ctx, user)
	if err != nil {
		return user, errorx.Internal("Can't create user")
	}

	return user, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (User, error) {
	user, err := s.store.GetByEmail(ctx, email)
	if err != nil {
		return user, errorx.NotFound("user not found")
	}

	return user, nil
}

func (s *userService) GetByID(ctx context.Context, id id.ID) (User, error) {
	user, err := s.store.GetByID(ctx, id)
	if err != nil {
		return user, errorx.NotFound("user not found")
	}

	return user, nil
}
