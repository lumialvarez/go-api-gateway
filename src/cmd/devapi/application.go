package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
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
	log.Print("Routes:")
	for _, routeItem := range routes {
		log.Print("--->> Path:", routeItem.RelativePath(), " -> To:", routeItem.UrlTarget(), " Type:", routeItem.TypeTarget())
	}

	r := gin.Default()

	r.GET("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, &routes) })
	r.POST("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, &routes) })
	r.PUT("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, &routes) })
	r.DELETE("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, &routes) })
	r.HEAD("/*proxyPath", func(ctx *gin.Context) { generic.HandlerManager(ctx, &routes) })

	//authSvc := *auth.RegisterRoutes(r, &config)
	//product.RegisterRoutes(r, &config, &authSvc)
	//order.RegisterRoutes(r, &config, &authSvc)
	r.Run(config.Port)

}
