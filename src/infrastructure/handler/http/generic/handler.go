package generic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/platform/httpclient"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"log"
	"net/http"
)

func Handler(ctx *gin.Context, route *domainRoute.Route) {
	path := ctx.Param("proxyPath")
	//method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", "*")
	fmt.Println(path)
	fmt.Println(route.UrlTarget() + path)

	err := httpclient.HttpPassThrough(route.UrlTarget(), ctx)
	if err != nil {
		log.Print("Error al invocar al servicio ", route.UrlTarget()+path, err)
		ctx.AbortWithStatus(http.StatusBadGateway)
	}

	/*
		if isPathMatch(path, "/api/gateway/v1/conf") {
			if method == "GET" {
				validRoute = true
				routeRepository := route.Init(config)
				useCaseGetRoute := get.NewUseCaseGetRoute(&routeRepository)
				var err error
				*routes, err = useCaseGetRoute.Execute()
				if err != nil {
					log.Fatalln("Failed to load routes", err)
				}
				log.Print("Routes:")
				for _, routeItem := range *routes {
					log.Print("--->> Path:", routeItem.RelativePath(), " -> To:", routeItem.UrlTarget(), " Type:", routeItem.TypeTarget())
				}
			}
		}*/
}

func RegisterHttpRoutes(r *gin.Engine, routes *[]domainRoute.Route) {
	for idx, _ := range *routes {
		routeItem := &(*routes)[idx]
		if routeItem.TypeTarget() == "http" {
			log.Print("Routes HTTP:")
			log.Print("--->> Path:", routeItem.RelativePath(), " -> To:", routeItem.UrlTarget(), " Type:", routeItem.TypeTarget())

			genericHandler := func(ctx *gin.Context) { Handler(ctx, routeItem) }

			r.Any(routeItem.RelativePath()+"/*proxyPath", genericHandler)

		}
	}
}
