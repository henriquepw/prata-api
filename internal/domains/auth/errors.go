package auth

import (
	"net/http"

	apiError "github.com/henriquepw/pobrin-api/pkg/errors"
)

var (
	ErrInvalidPassword = apiError.ServerError{
		Errors:     map[string]string{"Invalid user or password": "user and/or password is invalid"},
		Message:    "verify if user and password is correct",
		StatusCode: http.StatusUnauthorized,
	}
)
