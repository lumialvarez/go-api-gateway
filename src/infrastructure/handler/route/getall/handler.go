package getallRoutes

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall/mapper"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
	"net/http"
)

type UseCase interface {
	Execute() (*[]route.Route, error)
}

type ApiResponseProvider interface {
	ToAPIResponse(err error, cause ...string) *apierrors.APIError
}

type Handler struct {
	mapper              mapper.Mapper
	useCase             UseCase
	apiResponseProvider ApiResponseProvider
}

func NewHandler(useCase UseCase, apiResponseProvider ApiResponseProvider) Handler {
	return Handler{useCase: useCase, apiResponseProvider: apiResponseProvider}
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
