package devapi

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"log"
)

func Start() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	//authSvc := *auth.RegisterRoutes(r, &c)
	//product.RegisterRoutes(r, &c, &authSvc)
	//order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
