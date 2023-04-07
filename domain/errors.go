package domain

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	StatusCode int
	Message    string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("status %d: %s", e.StatusCode, e.Message)
}

var (
	ErrMessageNotFound CustomError = CustomError{
		StatusCode: http.StatusNotFound,
		Message:    "message not found",
	}
	ErrAuthorNotFound CustomError = CustomError{
		StatusCode: http.StatusNotFound,
		Message:    "author not found",
	}
)
