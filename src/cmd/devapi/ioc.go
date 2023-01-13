package devapi

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler"
	handlerGenericAuth "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/generic"
	handlerErrors "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/error"
	handlerGenericProfile "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/profile/generic"
	getallRoutes "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall"
	handlerRouteReload "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/reload"
	handlerSaveRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save"
	handlerUpdateRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/update"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/grpc/auth"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/grpc/profile"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/prometheus"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/provider/authorization"
	useCaseGetRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/getall"
	reloadRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/reload"
	saveRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/save"
	updateRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/update"
)

type DependenciesContainer struct {
	AuthorizationMiddleware authorization.Authentication
	PrometheusMiddleware    prometheus.PrometheusProvider
	Routes                  Routes
	Auth                    Auth
	Profile                 Profile
}

type Routes struct {
	GetAllRoutes handler.Handler
	ReloadRoutes handler.Routes
	SaveRoute    handler.Handler
	UpdateRoute  handler.Handler
}

type Auth struct {
	Generic handler.HandlerParams
}

type Profile struct {
	Generic handler.HandlerParams
}

func LoadDependencies(config config.Config) DependenciesContainer {
	repositoryRoutes := route.Init(config)
	apiProvider := handlerErrors.NewAPIResponseProvider()
	authServiceClient := auth.InitServiceClient(&config)
	profileServiceClient := profile.InitServiceClient(&config)
	authenticationService := authorization.NewAuthenticationService(authServiceClient)
	prometheusProvider := prometheus.NewPrometheusProvider()
	return DependenciesContainer{
		AuthorizationMiddleware: authenticationService,
		PrometheusMiddleware:    prometheusProvider,
		Routes: Routes{
			GetAllRoutes: newGetAllRoutesHandler(apiProvider, repositoryRoutes),
			ReloadRoutes: newReloadRoutesHandler(apiProvider, repositoryRoutes),
			SaveRoute:    newSaveRouteHandler(apiProvider, repositoryRoutes),
			UpdateRoute:  newUpdateRouteHandler(apiProvider, repositoryRoutes),
		},
		Auth: Auth{
			Generic: handlerGenericAuth.NewHandler(apiProvider, handlerGenericAuth.Authentication{AuthServiceClient: authServiceClient}),
		},
		Profile: Profile{
			Generic: handlerGenericProfile.NewHandler(apiProvider, handlerGenericProfile.ProfileService{ProfileServiceClient: profileServiceClient}),
		},
	}
}

func newGetAllRoutesHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository) handler.Handler {
	useCase := useCaseGetRoute.NewUseCaseGetRoute(&repository)

	return getallRoutes.NewHandler(useCase, apiProvider)
}

func newReloadRoutesHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository) handler.Routes {
	useCaseGetRoute := reloadRoute.NewUseCaseReloadRoute(&repository)

	return handlerRouteReload.NewHandler(useCaseGetRoute, apiProvider)
}

func newSaveRouteHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository) handler.Handler {
	useCase := saveRoute.NewUseCaseSaveRoute(&repository)

	return handlerSaveRoute.NewHandler(useCase, apiProvider)
}

func newUpdateRouteHandler(apiProvider *handlerErrors.APIResponseProvider, repository route.Repository) handler.Handler {
	useCase := updateRoute.NewUseCaseUpdateRoute(&repository)

	return handlerUpdateRoute.NewHandler(useCase, apiProvider)
}
