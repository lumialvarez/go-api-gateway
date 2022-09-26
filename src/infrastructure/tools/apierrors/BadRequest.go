package apierrors

import "net/http"

func NewBadRequestError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusBadRequest,
		Message: message,
		Err:     "Bad Request Error",
		Cause:   cause,
	}
}
