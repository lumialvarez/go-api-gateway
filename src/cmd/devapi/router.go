package devapi

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/http/generic"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ConfigureRoutes(r *gin.Engine, config config.Config, dynamicRoutes *[]route.Route) {
	handlers := LoadDependencies(config)

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	// Prometheus endpoint
	r.GET("/prometheus", gin.WrapH(promhttp.Handler()))

	r.Use(handlers.PrometheusMiddleware.PrometheusMiddleware)

	//Map all http dynamic Routes
	generic.RegisterHttpRoutes(r, handlers.AuthorizationMiddleware.AuthRequired, dynamicRoutes)

	//Map internal and provider services
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
	authGroupInternal.POST("/user", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "register") })
	authGroupInternal.GET("/user", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "list") })
	authGroupInternal.PUT("/user", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "update") })

	authGroupInternal.GET("/user/current/notification", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "current") })
	authGroupInternal.PUT("/user/current/notification/:notificationId", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "read_notification") })

	authGroupExternal := authGroup.Group("/ext")
	authGroupExternal.POST("/user/validate", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "validate") })
	authGroupExternal.POST("/user/login", func(ctx *gin.Context) { handlers.Auth.Generic.Handler(ctx, "login") })

	//API Profile methods
	profileGroup := r.Group("/profile/api/v1")
	profileGroupInternal := profileGroup.Group("/int")
	profileGroupInternal.Use(handlers.AuthorizationMiddleware.AuthRequired)
	profileGroupInternal.GET("/profile", func(ctx *gin.Context) { handlers.Profile.Generic.Handler(ctx, "list") })
	profileGroupInternal.POST("/profile", func(ctx *gin.Context) { handlers.Profile.Generic.Handler(ctx, "save") })
	profileGroupInternal.PUT("/profile", func(ctx *gin.Context) { handlers.Profile.Generic.Handler(ctx, "update") })

	profileGroupExternal := profileGroup.Group("/ext")
	profileGroupExternal.GET("/profile", func(ctx *gin.Context) { handlers.Profile.Generic.Handler(ctx, "list") })

}
