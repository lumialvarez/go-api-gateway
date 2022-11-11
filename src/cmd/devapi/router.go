package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/http/generic"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

func ConfigureRoutes(r *gin.Engine, config config.Config, dynamicRoutes *[]route.Route) {
	handlers := LoadDependencies(config)

	//Map all http dynamic Routes
	generic.RegisterHttpRoutes(r, handlers.AuthorizationMiddleware.AuthRequired, dynamicRoutes)

	registerEndpoints(r, handlers, dynamicRoutes)
}

func registerEndpoints(r *gin.Engine, handlers DependenciesContainer, dynamicRoutes *[]route.Route) {

	//API gateway methods
	gatewayGroup := r.Group("/gateway")
	gatewayGroup.Use(handlers.AuthorizationMiddleware.AuthRequiredAsAdmin)
	gatewayGroup.GET("/api/v1/int/conf/route", handlers.Routes.GetAllRoutes.Handler)
	gatewayGroup.POST("/api/v1/int/conf/route", handlers.Routes.SaveRoute.Handler)
	gatewayGroup.PUT("/api/v1/int/conf/route", handlers.Routes.UpdateRoute.Handler)
	gatewayGroup.POST("/api/v1/int/conf/route/reload", func(ctx *gin.Context) { handlers.Routes.ReloadRoutes.Handler(ctx, dynamicRoutes) })

	//API Authorization methods
	authGroup := r.Group("/authorization/api/v2")
	authGroupInternal := authGroup.Group("/int")
	authGroupInternal.Use(handlers.AuthorizationMiddleware.AuthRequiredAsAdmin)

	authGroupExternal := authGroup.Group("/ext")
	authGroupExternal.POST("/auth/validate", handlers.Auth.Validate.Handler)
}
