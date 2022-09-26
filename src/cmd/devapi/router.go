package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/http/generic"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

func ConfigureRoutes(r *gin.Engine, config config.Config, dynamicRoutes *[]route.Route) {

	//Map all http dynamic Routes
	generic.RegisterHttpRoutes(r, dynamicRoutes)

	handlers := LoadDependencies(config)

	registerEndpoints(r, handlers, dynamicRoutes)

	//authSvc := *auth.ConfigureRoutes(r, &config)
	//product.ConfigureRoutes(r, &config, &authSvc)
}

func registerEndpoints(r *gin.Engine, handlers DependenciesContainer, dynamicRoutes *[]route.Route) {

	//API gateway methods
	gatewayGroup := r.Group("/gateway")
	gatewayGroup.GET("/api/v1/conf/route", handlers.GetRoutes.Handler)
	gatewayGroup.POST("/api/v1/conf/route/reload", func(ctx *gin.Context) { handlers.ReloadRoutes.Handler(ctx, dynamicRoutes) })

}
