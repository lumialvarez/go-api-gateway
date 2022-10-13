package apierrors

import "net/http"

func NewUnauthorizedError(message string, cause ...string) *APIError {
	return &APIError{
		Status:  http.StatusUnauthorized,
		Message: message,
		Err:     "Not Authorized",
		Cause:   cause,
	}
}
