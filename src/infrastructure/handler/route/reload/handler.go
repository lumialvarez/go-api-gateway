package handlerRouteReload

import (
	"github.com/gin-gonic/gin"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"log"
	"net/http"
)

type UseCase interface {
	Execute(ctx gin.Context, routes *[]domainRoute.Route) ([]domainRoute.Route, error)
}

type ApiResponseProvider interface {
	ToAPIResponse(err error, cause ...string) *apierrors.APIError
}

type AuthenticationService interface {
	IsAuthorized(ctx *gin.Context) (bool, error)
}

type Handler struct {
	useCase               UseCase
	apiResponseProvider   ApiResponseProvider
	authenticationService AuthenticationService
}

func NewHandler(useCase UseCase, apiResponseProvider ApiResponseProvider, authenticationService AuthenticationService) Handler {
	return Handler{useCase: useCase, apiResponseProvider: apiResponseProvider, authenticationService: authenticationService}
}

func (h Handler) Handler(ctx *gin.Context, routes *[]domainRoute.Route) {
	if authorized, err := h.authenticationService.IsAuthorized(ctx); !authorized || err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	_, err := h.useCase.Execute(*ctx, routes)

	for _, routeItem := range *routes {
		log.Print("Reload target url in routes HTTP:")
		log.Print("--->> Path:", routeItem.RelativePath(), " -> To:", routeItem.UrlTarget(), " Type:", routeItem.TypeTarget())
	}

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
