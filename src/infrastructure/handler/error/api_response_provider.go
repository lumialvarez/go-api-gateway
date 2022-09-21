package error

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/tools/apierror"
	"net/http"
)

const (
	badRequest    string = "bad request"
	internalError string = "internal error"
)

type APIResponseProvider struct {
}

func NewAPIResponseProvider() *APIResponseProvider {
	return &APIResponseProvider{}
}

func (a *APIResponseProvider) NewBadRequest(message string, cause ...string) *apierror.APIError {
	return &apierror.APIError{
		Status:  http.StatusBadRequest,
		Message: message,
		Err:     badRequest,
		Cause:   cause,
	}
}

func (a *APIResponseProvider) NewInternalServerError(message string, cause ...string) *apierror.APIError {
	return &apierror.APIError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Err:     internalError,
		Cause:   cause,
	}
}

func (a *APIResponseProvider) ToAPIResponse(err error, cause ...string) *apierror.APIError {
	switch err.(type) {
	case apierror.BadRequest:
		return a.NewBadRequest(err.Error(), cause...)
	default:
		return a.NewInternalServerError(err.Error())
	}
}
