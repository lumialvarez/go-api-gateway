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
	ctx.Header("Access-Control-Allow-Origin", "*")
	fmt.Println(path)
	fmt.Println(route.UrlTarget() + path)

	if !route.Enable() {
		log.Println("Endpoint inhabilitado ", route.UrlTarget()+path)
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	err := httpclient.HttpPassThrough(route.UrlTarget(), ctx)
	if err != nil {
		log.Println("Error al invocar al servicio ", route.UrlTarget()+path, err)
		ctx.AbortWithStatus(http.StatusBadGateway)
	}
}

func RegisterHttpRoutes(r *gin.Engine, authorizationFunction gin.HandlerFunc, routes *[]domainRoute.Route) {
	for idx, _ := range *routes {
		routeItem := &(*routes)[idx]
		if routeItem.TypeTarget() == "http" && routeItem.Enable() {
			log.Print("Routes HTTP:")
			log.Print("--->> Path:", routeItem.RelativePath(), " -> To:", routeItem.UrlTarget(), " Type:", routeItem.TypeTarget())

			genericHandler := func(ctx *gin.Context) { Handler(ctx, routeItem) }
			genericGroup := r.Group("/" + routeItem.RelativePath())
			if routeItem.Secure() {
				genericGroup.Use(authorizationFunction)
			}
			genericGroup.Any("/*proxyPath", genericHandler)

		}
	}
}
