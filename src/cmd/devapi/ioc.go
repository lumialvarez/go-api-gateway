package devapi

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler"
	handlerErrors "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/error"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/get"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/get/mapper"
	handlerRouteReload "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/reload"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	useCaseGetRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/get"
	updateRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/update"
)

type DependenciesContainer struct {
	GetRoutes    handler.Handler
	ReloadRoutes handler.Routes
}

func LoadDependencies(config config.Config) DependenciesContainer {
	repositoryRoutes := route.Init(config)
	apiProvider := handlerErrors.NewAPIResponseProvider()
	return DependenciesContainer{
		GetRoutes:    newGetRoutesHandler(apiProvider, repositoryRoutes),
		ReloadRoutes: newReloadRoutesHandler(apiProvider, repositoryRoutes),
	}
}

func newGetRoutesHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository) handler.Handler {
	useCaseGetRoute := useCaseGetRoute.NewUseCaseGetRoute(&repository)
	mapperRoute := mapper.Mapper{}

	return get.NewHandler(mapperRoute, useCaseGetRoute, apiProvider)
}

func newReloadRoutesHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository) handler.Routes {
	useCaseGetRoute := updateRoute.NewUseCaseGetRoute(&repository)

	return handlerRouteReload.NewHandler(useCaseGetRoute, apiProvider)
}
