package generic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/platform/httpclient"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"log"
	"net/http"
	"strings"
)

func HandlerManager(ctx *gin.Context, routes *[]route.Route) {
	path := ctx.Param("proxyPath")
	fmt.Println(path)

	validRoute := false
	for _, routeItem := range *routes {
		fmt.Println(routeItem)
		if routeItem.TypeTarget() == "http" {
			if isPathMatch(path, routeItem.RelativePath()) {
				validRoute = true
				err := httpclient.HttpPassThrough(routeItem.UrlTarget(), ctx)
				if err != nil {
					log.Fatal("Error al invocar al servicio ", routeItem.UrlTarget()+path, err)
					ctx.AbortWithStatus(http.StatusBadGateway)
				}
			}
		}
	}

	if !validRoute {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func isPathMatch(path string, pathCandidate string) bool {
	//Fixme mejorar match
	return strings.Contains(path, pathCandidate)
}
