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
	msg := "Requisição inválida"
	if message != nil {
		msg = strings.Join(message, "")
	}

	return ServerError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
	}
}

func NotFound(message ...string) ServerError {
	msg := "Dado não encontrado"
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
		Message:    "Não Autorizado",
		StatusCode: http.StatusUnauthorized,
	}
}

func Conflict(message ...string) ServerError {
	msg := "Conflito"
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
		Message:    "Erro de validação",
		Errors:     errors,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func Internal(msgs ...string) ServerError {
	message := "Error Interno"
	if msgs != nil {
		message = strings.Join(msgs, "; ")
	}

	return ServerError{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func InvalidJSON(msgs ...string) ServerError {
	message := "Dado mal formatado"
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
		Message:    "Método inválido",
		StatusCode: http.StatusMethodNotAllowed,
	}
}

func UnprocessableEntity(message ...string) ServerError {
	msg := "Requisição não pode ser processada"
	if message != nil {
		msg = strings.Join(message, "")
	}

	return ServerError{
		Message:    msg,
		StatusCode: http.StatusUnprocessableEntity,
	}
}
