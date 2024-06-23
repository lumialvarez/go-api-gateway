package generic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/platform/httpclient"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"log"
	"net/http"
	"strings"
)

const (
	anyRelativePath = "/*proxyPath"
)

func Handler(ctx *gin.Context, route *domainRoute.Route) {
	path := ctx.Param("proxyPath")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", strings.Join(route.GetStringMethods(), ", "))
	fmt.Println(path)
	fmt.Println(route.UrlTarget() + path)

	if !route.Enable() {
		log.Println("disabled Endpoint ", route.UrlTarget()+path)
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	err := httpclient.HttpPassThrough(route.UrlTarget(), ctx)
	if err != nil {
		log.Println("error invoking service ", route.UrlTarget()+path, err)
		ctx.AbortWithStatus(http.StatusBadGateway)
	}
}

func RegisterHttpRoutes(r *gin.Engine, authorizationFunction gin.HandlerFunc, routes *[]domainRoute.Route) {
	for idx := range *routes {
		routeItem := &(*routes)[idx]
		if routeItem.TypeTarget() == "http" && routeItem.Enable() {
			log.Print("Routes HTTP:")
			log.Print("--->> Path:", routeItem.RelativePath(), " -> To:", routeItem.UrlTarget(), " Type:", routeItem.TypeTarget(), " Methods:", strings.Join(routeItem.GetStringMethods(), ","))

			genericHandler := func(ctx *gin.Context) { Handler(ctx, routeItem) }
			genericGroup := r.Group("/" + routeItem.RelativePath())
			if routeItem.Secure() {
				genericGroup.Use(authorizationFunction)
			}
			genericGroup.Match(routeItem.GetStringMethods(), anyRelativePath, genericHandler)

		}
	}
}
