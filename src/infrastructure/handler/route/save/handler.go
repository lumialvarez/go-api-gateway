package handlerSaveRoute

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save/contract"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save/mapper"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
)

const (
	invalidFormat string = "invalid_message_format"
)

type UseCase interface {
	Execute(route route.Route) error
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
	var request contract.SaveRouteRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	domainRoute := h.mapper.ToDomain(request)

	err := h.useCase.Execute(domainRoute)
	if err != nil {
		return h.apiResponseProvider.ToAPIResponse(err)
	}
	return nil
}
