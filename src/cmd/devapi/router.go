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
	gatewayGroup := r.Group("/gateway/api/v1")
	gatewayGroup.Use(handlers.AuthorizationMiddleware.AuthRequiredAsAdmin)
	gatewayGroup.GET("/int/conf/route", handlers.Routes.GetAllRoutes.Handler)
	gatewayGroup.POST("/int/conf/route", handlers.Routes.SaveRoute.Handler)
	gatewayGroup.PUT("/int/conf/route", handlers.Routes.UpdateRoute.Handler)
	gatewayGroup.POST("/int/conf/route/reload", func(ctx *gin.Context) { handlers.Routes.ReloadRoutes.Handler(ctx, dynamicRoutes) })

	//API Authorization methods
	authGroup := r.Group("/authorization/api/v1")
	authGroupInternal := authGroup.Group("/int")
	authGroupInternal.Use(handlers.AuthorizationMiddleware.AuthRequiredAsAdmin)
	authGroupInternal.POST("/auth/user", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "register") })
	authGroupInternal.GET("/auth/user", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "list") })
	authGroupInternal.PUT("/auth/user", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "update") })

	authGroupExternal := authGroup.Group("/ext")
	authGroupExternal.POST("/auth/validate", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "validate") })
	authGroupExternal.POST("/auth/login", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "login") })
}
