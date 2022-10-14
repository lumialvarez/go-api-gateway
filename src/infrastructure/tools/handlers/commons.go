package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
)

// HandlerFunc is the func type for the custom handlers.
type HandlerFunc func(c *gin.Context) *apierrors.APIError

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func ErrorWrapper(handlerFunc HandlerFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}
