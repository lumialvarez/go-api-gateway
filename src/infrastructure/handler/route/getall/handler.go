package getallRoutes

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall/contract"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/tools/handlers"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
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

type AuthenticationService interface {
	IsAuthorized(ctx *gin.Context) (bool, error)
}

type Handler struct {
	mapper                Mapper
	useCase               UseCase
	apiResponseProvider   ApiResponseProvider
	authenticationService AuthenticationService
}

func NewHandler(mapper Mapper, useCase UseCase, apiResponseProvider ApiResponseProvider, authenticationService AuthenticationService) Handler {
	return Handler{mapper: mapper, useCase: useCase, apiResponseProvider: apiResponseProvider, authenticationService: authenticationService}
}

func (h Handler) Handler(ginCtx *gin.Context) {
	handlers.ErrorWrapper(h.handler, ginCtx)
}

func (h Handler) handler(ctx *gin.Context) *apierrors.APIError {
	if authorized, err := h.authenticationService.IsAuthorized(ctx); !authorized || err != nil {
		return apierrors.NewUnauthorizedError("Not Authorized")
	}
	domainRoutes, err := h.useCase.Execute()
	if err != nil {
		return h.apiResponseProvider.ToAPIResponse(err)
	}
	dtoRoutes := h.mapper.ToDTOs(*domainRoutes)
	ctx.JSON(http.StatusOK, dtoRoutes)
	return nil
}
