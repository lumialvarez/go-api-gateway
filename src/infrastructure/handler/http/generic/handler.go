package generic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/platform/httpclient"
	"log"
	"net/http"
	"strings"
)

func HandlerManager(ctx *gin.Context, c config.Config) {
	path := ctx.Param("proxyPath")
	fmt.Println(path)

	if isPathMatch(path, "api/ext/v1") {
		err := httpclient.HttpPassThrough(c.PersonalWebsiteServices, ctx)
		if err != nil {
			log.Fatal("Error al invocar al servicio ", c.PersonalWebsiteServices+path, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func isPathMatch(path string, pathCandidate string) bool {
	//Fixme mejorar match
	return strings.Contains(path, pathCandidate)
}
