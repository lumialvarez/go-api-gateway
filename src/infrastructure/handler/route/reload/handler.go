package handlerRouteReload

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/tools/apierrors"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"net/http"
)

type UseCase interface {
	Execute(ctx gin.Context, routes *[]domainRoute.Route) ([]domainRoute.Route, error)
}

type ApiResponseProvider interface {
	ToAPIResponse(err error, cause ...string) *apierrors.APIError
}

type Handler struct {
	useCase             UseCase
	apiResponseProvider ApiResponseProvider
}

func NewHandler(useCase UseCase, apiResponseProvider ApiResponseProvider) Handler {
	return Handler{useCase: useCase, apiResponseProvider: apiResponseProvider}
}

func (h Handler) Handler(ctx *gin.Context, routes *[]domainRoute.Route) {
	_, err := h.useCase.Execute(*ctx, routes)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
