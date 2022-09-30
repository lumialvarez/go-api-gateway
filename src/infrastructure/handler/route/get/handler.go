package get

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/get/contract"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/tools/apierrors"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/tools/handlers"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"net/http"
)

type Mapper interface {
	ToDTOs(domainRoutes []route.Route) []contract.GetRouteResponse
}

type UseCase interface {
	Execute() (*[]route.Route, error)
}

type ApiResponseProvider interface {
	ToAPIResponse(err error, cause ...string) *apierrors.APIError
}

type Handler struct {
	mapper              Mapper
	useCase             UseCase
	apiResponseProvider ApiResponseProvider
}

func NewHandler(mapper Mapper, useCase UseCase, apiResponseProvider ApiResponseProvider) Handler {
	return Handler{mapper: mapper, useCase: useCase, apiResponseProvider: apiResponseProvider}
}

func (h Handler) Handler(ginCtx *gin.Context) {
	handlers.ErrorWrapper(h.handler, ginCtx)
}

func (h Handler) handler(ctx *gin.Context) *apierrors.APIError {

	domainRoutes, err := h.useCase.Execute()
	if err != nil {
		return h.apiResponseProvider.ToAPIResponse(err)

	}
	dtoRoutes := h.mapper.ToDTOs(*domainRoutes)
	ctx.JSON(http.StatusOK, dtoRoutes)
	return nil
}
