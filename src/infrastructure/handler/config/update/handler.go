package update

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/get"
	"log"
)

func Handler(r *gin.Engine, routes *[]domainRoute.Route, config config.Config) {
	//Map all http routes
	//generic.RegisterHttpRoutes(r, routes)

	routeRepository := route.Init(config)
	useCaseGetRoute := get.NewUseCaseGetRoute(&routeRepository)
	var err error
	err = useCaseGetRoute.Update(routes)
	if err != nil {
		log.Fatalln("Failed to load routes", err)
	}

	//Map all http routes
	//generic.RegisterHttpRoutes(r, routes)
}
