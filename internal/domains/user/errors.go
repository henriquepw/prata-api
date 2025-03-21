package user

import (
	"net/http"

	apiError "github.com/henriquepw/pobrin-api/pkg/errors"
)

var (
	ErrUserNotFound = apiError.ServerError{
		Errors:     map[string]string{"Not Found": "user not found"},
		Message:    "could not found user in database, verify if request data is correct",
		StatusCode: http.StatusNotFound,
	}
)
