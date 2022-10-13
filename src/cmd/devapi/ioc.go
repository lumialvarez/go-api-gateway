package devapi

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler"
	handlerErrors "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/error"
	getallRoutes "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall/mapper"
	handlerRouteReload "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/reload"
	handlerSaveRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save"
	mapperSaveRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save/mapper"
	handlerUpdateRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/update"
	mapperUpdateRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/update/mapper"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/authentication"
	useCaseGetRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/getall"
	reloadRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/reload"
	saveRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/save"
	updateRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/update"
)

type DependenciesContainer struct {
	GetAllRoutes handler.Handler
	ReloadRoutes handler.Routes
	SaveRoute    handler.Handler
	UpdateRoute  handler.Handler
}

func LoadDependencies(config config.Config) DependenciesContainer {
	repositoryRoutes := route.Init(config)
	apiProvider := handlerErrors.NewAPIResponseProvider()
	authenticationServiceAdmin := authentication.NewAuthenticationService(authentication.Admin)
	return DependenciesContainer{
		GetAllRoutes: newGetAllRoutesHandler(apiProvider, repositoryRoutes, authenticationServiceAdmin),
		ReloadRoutes: newReloadRoutesHandler(apiProvider, repositoryRoutes, authenticationServiceAdmin),
		SaveRoute:    newSaveRouteHandler(apiProvider, repositoryRoutes, authenticationServiceAdmin),
		UpdateRoute:  newUpdateRouteHandler(apiProvider, repositoryRoutes, authenticationServiceAdmin),
	}
}

func newGetAllRoutesHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository, authenticationService authentication.Authentication) handler.Handler {
	useCase := useCaseGetRoute.NewUseCaseGetRoute(&repository)
	mapper := mapperGetAllRoutes.Mapper{}

	return getallRoutes.NewHandler(mapper, useCase, apiProvider, authenticationService)
}

func newReloadRoutesHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository, authenticationService authentication.Authentication) handler.Routes {
	useCaseGetRoute := reloadRoute.NewUseCaseReloadRoute(&repository)

	return handlerRouteReload.NewHandler(useCaseGetRoute, apiProvider, authenticationService)
}

func newSaveRouteHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository, authenticationService authentication.Authentication) handler.Handler {
	useCase := saveRoute.NewUseCaseSaveRoute(&repository)
	mapper := mapperSaveRoute.Mapper{}

	return handlerSaveRoute.NewHandler(mapper, useCase, apiProvider, authenticationService)
}

func newUpdateRouteHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository, authenticationService authentication.Authentication) handler.Handler {
	useCase := updateRoute.NewUseCaseUpdateRoute(&repository)
	mapper := mapperUpdateRoute.Mapper{}

	return handlerUpdateRoute.NewHandler(mapper, useCase, apiProvider, authenticationService)
}
