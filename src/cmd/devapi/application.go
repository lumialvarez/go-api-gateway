package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/config/update"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/http/generic"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	"github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/get"
	"log"
)

func Start() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	routeRepository := route.Init(config)
	useCaseGetRoute := get.NewUseCaseGetRoute(&routeRepository)
	routes, err := useCaseGetRoute.Execute()
	if err != nil {
		log.Fatalln("Failed to load routes", err)
	}

	r := gin.Default()

	//Map all http routes
	generic.RegisterHttpRoutes(r, routes)

	r.GET("/gateway/v1/conf", func(ctx *gin.Context) { update.Handler(r, routes, config) })

	//authSvc := *auth.RegisterRoutes(r, &config)
	//product.RegisterRoutes(r, &config, &authSvc)
	r.Run(config.Port)
}
