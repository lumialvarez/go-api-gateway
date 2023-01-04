package devapi

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler"
	handlerListAuth "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/list"
	handlerLoginAuth "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/login"
	handlerRegisterAuth "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/register"
	handlerUpdateAuth "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/update"
	handlerValidateAuth "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/validate"
	handlerErrors "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/error"
	getallRoutes "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall"
	handlerRouteReload "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/reload"
	handlerSaveRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save"
	handlerUpdateRoute "github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/update"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/authentication"
	useCaseGetRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/getall"
	reloadRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/reload"
	saveRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/save"
	updateRoute "github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/update"
)

type DependenciesContainer struct {
	AuthorizationMiddleware authentication.Authentication
	Routes                  Routes
	Auth                    Auth
}

type Routes struct {
	GetAllRoutes handler.Handler
	ReloadRoutes handler.Routes
	SaveRoute    handler.Handler
	UpdateRoute  handler.Handler
}

type Auth struct {
	Login    handler.Handler
	Register handler.Handler
	Validate handler.Handler
	List     handler.Handler
	Update   handler.Handler
}

func LoadDependencies(config config.Config) DependenciesContainer {
	repositoryRoutes := route.Init(config)
	apiProvider := handlerErrors.NewAPIResponseProvider()
	authServiceClient := authentication.InitServiceClient(&config)
	authenticationService := authentication.NewAuthenticationService(authServiceClient)
	return DependenciesContainer{
		AuthorizationMiddleware: authenticationService,
		Routes: Routes{
			GetAllRoutes: newGetAllRoutesHandler(apiProvider, repositoryRoutes),
			ReloadRoutes: newReloadRoutesHandler(apiProvider, repositoryRoutes),
			SaveRoute:    newSaveRouteHandler(apiProvider, repositoryRoutes),
			UpdateRoute:  newUpdateRouteHandler(apiProvider, repositoryRoutes),
		},
		Auth: Auth{
			Login:    handlerLoginAuth.NewHandler(apiProvider, handlerLoginAuth.Authentication{AuthServiceClient: authServiceClient}),
			Register: handlerRegisterAuth.NewHandler(apiProvider, handlerRegisterAuth.Authentication{AuthServiceClient: authServiceClient}),
			Validate: handlerValidateAuth.NewHandler(apiProvider, handlerValidateAuth.Authentication{AuthServiceClient: authServiceClient}),
			List:     handlerListAuth.NewHandler(apiProvider, handlerListAuth.Authentication{AuthServiceClient: authServiceClient}),
			Update:   handlerUpdateAuth.NewHandler(apiProvider, handlerUpdateAuth.Authentication{AuthServiceClient: authServiceClient}),
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
