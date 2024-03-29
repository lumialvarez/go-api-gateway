package handler

import (
	"github.com/gin-gonic/gin"
	domainRoute "github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Handler interface {
	Handler(ginCtx *gin.Context)
}

type HandlerParams interface {
	Handler(ginCtx *gin.Context, params ...string)
}

type Routes interface {
	Handler(ctx *gin.Context, routes *[]domainRoute.Route)
}
