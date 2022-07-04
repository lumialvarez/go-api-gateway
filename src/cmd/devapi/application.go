package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/http/generic"
	"log"
)

func Start() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	r.GET("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, c) })
	r.POST("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, c) })
	r.PUT("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, c) })
	r.DELETE("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, c) })
	r.HEAD("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, c) })

	//authSvc := *auth.RegisterRoutes(r, &c)
	//product.RegisterRoutes(r, &c, &authSvc)
	//order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)

}
