package dto

import (
	"database/sql"
	"errors"
	"net/http"
)

type HTTPError struct {
	Code    int
	Message string
}

var (
	ErrResourceNotFound = &HTTPError{
		Code:    http.StatusNotFound,
		Message: "the requested resource could not be found",
	}
)

func FromError(err error) *HTTPError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrResourceNotFound
	default:
		return &HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong",
		}
	}
}
