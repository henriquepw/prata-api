package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type ServerError struct {
	Errors     map[string]string `json:"errors,omitempty"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
}

func (s ServerError) Error() string {
	return fmt.Sprintf("SERVER_ERROR [%v]: %v.\n%v\n--", s.StatusCode, s.Message, s.Errors)
}

func BadRequest(message ...string) ServerError {
	msg := "Bad Request"
	if message != nil {
		msg = strings.Join(message, "")
	}

	return ServerError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
	}
}

func NotFound(message ...string) ServerError {
	msg := "Not found"
	if message != nil {
		msg = strings.Join(message, "")
	}

	return ServerError{
		Message:    msg,
		StatusCode: http.StatusNotFound,
	}
}

func Unauthorized() ServerError {
	return ServerError{
		Message:    "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func Conflict(message ...string) ServerError {
	msg := "Conflict"
	if message != nil {
		msg = strings.Join(message, "")
	}

	return ServerError{
		Message:    msg,
		StatusCode: http.StatusConflict,
	}
}

func InvalidRequestData(errors map[string]string) ServerError {
	return ServerError{
		Message:    "Validation error",
		Errors:     errors,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func Internal(msgs ...string) ServerError {
	message := "Internal error"
	if msgs != nil {
		message = strings.Join(msgs, "; ")
	}

	return ServerError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func InvalidJSON(msgs ...string) ServerError {
	message := "Invalid data format"
	if msgs != nil {
		message = strings.Join(msgs, "; ")
	}

	return ServerError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func MethodNotAllowed() ServerError {
	return ServerError{
		Message:    "Method not allowed",
		StatusCode: http.StatusMethodNotAllowed,
	}
}
