// Package web implements functions for web rest api
package web

import (
	"context"
	"net/http"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error
