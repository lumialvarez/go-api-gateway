package generic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/platform/httpclient"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-api-gateway/src/internal/route/usecase/get"
	"log"
	"net/http"
	"strings"
)

func HandlerManager(ctx *gin.Context, routes *[]domainRoute.Route, config config.Config) {
	path := ctx.Param("proxyPath")
	method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", "*")
	fmt.Println(path)

	validRoute := false
	for _, routeItem := range *routes {
		fmt.Println(routeItem)
		if routeItem.TypeTarget() == "http" {
			if isPathMatch(path, routeItem.RelativePath()) {
				validRoute = true
				err := httpclient.HttpPassThrough(routeItem.UrlTarget(), ctx)
				if err != nil {
					log.Print("Error al invocar al servicio ", routeItem.UrlTarget()+path, err)
					ctx.AbortWithStatus(http.StatusBadGateway)
				}
			}
		}
	}

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
	}

	if !validRoute {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func isPathMatch(path string, pathCandidate string) bool {
	//Fixme mejorar match
	return strings.Contains(path, pathCandidate)
}
