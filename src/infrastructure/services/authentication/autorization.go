package authentication

import (
	"github.com/gin-gonic/gin"
)

type Authentication struct {
	rol Rol
}

func NewAuthenticationService(rol Rol) Authentication {
	return Authentication{rol: rol}
}

func (auth Authentication) IsAuthorized(ctx *gin.Context) (bool, error) {
	//todo disable for go-grpc-auth-service integration
	return false, nil
}
