package service

import (
	"context"
	"github.com/henriquepw/pobrin-api/pkg/page"
)

type Service interface {
}

type CRUDService[T any, Q any, C any, U any, P any] interface {
	Service
	// List returns a paginated list of entities
	List(ctx context.Context, q Q) (*page.Page[T], error)
	// Create creates a new entity
	Create(ctx context.Context, d C) (*T, error)
	// Update updates an existing entity
	Update(ctx context.Context, id string, d U) (*T, error)
	// Patch updates an existing entity partially
	Patch(ctx context.Context, id string, d P) (*T, error)
	// Get returns an entity by id
	Get(ctx context.Context, id string) (*T, error)
	// Delete deletes an entity by id
	Delete(ctx context.Context, id string) error
}
