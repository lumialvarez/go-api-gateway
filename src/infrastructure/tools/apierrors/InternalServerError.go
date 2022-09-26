package apierrors

import "net/http"

func NewInternalServerError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusBadRequest,
		Message: message,
		Err:     "Internal Server Error",
		Cause:   cause,
	}
}
