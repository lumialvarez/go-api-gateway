package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/http/generic"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/reload"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

func RegisterRoutes(r *gin.Engine, config config.Config, dynamicRoutes *[]route.Route) {

	//Map all http dynamic Routes
	generic.RegisterHttpRoutes(r, dynamicRoutes)

	//API gateway methods
	gatewayGroup := r.Group("/gateway")
	gatewayGroup.POST("/api/v1/conf/reload", func(ctx *gin.Context) { reload.Handler(r, dynamicRoutes, config) })

	//authSvc := *auth.RegisterRoutes(r, &config)
	//product.RegisterRoutes(r, &config, &authSvc)
}
