package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/getEnabled"
	"log"
)

func Start() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	var dynamicRoutes *[]domainRoute.Route
	dynamicRoutes, err = loadRoutes(config)
	if err != nil {
		log.Fatalln("Failed to load dynamic routes", err)
	}

	r := gin.Default()

	ConfigureRoutes(r, config, dynamicRoutes)

	r.Run(config.Port)
}

func loadRoutes(config config.Config) (*[]domainRoute.Route, error) {
	routeRepository := route.Init(config)
	useCaseGetRoute := getEnabled.NewUseCaseGetEnabledRoute(&routeRepository)
	routes, err := useCaseGetRoute.Execute()
	return routes, err
}
