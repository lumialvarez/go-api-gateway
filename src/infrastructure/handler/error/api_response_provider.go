package handlerErrors

import (
	"github.com/lumialvarez/go-common-tools/web/apierrors"
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

/*
func (a *APIResponseProvider) NewBadRequest(message string, cause ...string) *apierror.APIError {
	return &apierror.APIError{
		Status:  http.StatusBadRequest,
		Message: message,
		Err:     badRequest,
		Cause:   cause,
	}
}
*/

func (a *APIResponseProvider) NewInternalServerError(message string, cause ...string) *apierrors.APIError {
	return apierrors.NewInternalServerError(message, cause...)
}

func (a *APIResponseProvider) ToAPIResponse(err error, cause ...string) *apierrors.APIError {
	switch err.(type) {
	/*case apierrors.BadRequest:
	return a.NewBadRequest(err.Error(), cause...)*/
	default:
		return a.NewInternalServerError(err.Error())
	}
}
