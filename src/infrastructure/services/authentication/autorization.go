package authentication

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

type Authentication struct {
	authServiceClient *ServiceClient
}

func NewAuthenticationService(authServiceClient *ServiceClient) Authentication {
	return Authentication{authServiceClient: authServiceClient}
}

func (auth *Authentication) AuthRequired(ctx *gin.Context) {
	_, err := auth.processAuthorization(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (auth *Authentication) AuthRequiredAsAdmin(ctx *gin.Context) {
	_, err := auth.processAuthorization(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (auth *Authentication) processAuthorization(ctx *gin.Context) (*pb.ValidateResponse, error) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		//ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, errors.New("Unauthorized")
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		//ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, errors.New("Unauthorized")
	}

	res, err := auth.authServiceClient.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		//ctx.AbortWithStatus(http.StatusUnauthorized)
		return nil, errors.New("Unauthorized")
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()

	return res, nil
}
