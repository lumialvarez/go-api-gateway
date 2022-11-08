package authentication

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"log"
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
	handlers.ErrorWrapper(auth.processAuthorization, ctx)
	if ctx.Writer.Status() != 200 && ctx.Writer.Status() != 201 {
		ctx.Abort()
	}
}

func (auth *Authentication) AuthRequiredAsAdmin(ctx *gin.Context) {
	handlers.ErrorWrapper(auth.processAuthorization, ctx)
	if ctx.Writer.Status() != 200 && ctx.Writer.Status() != 201 {
		ctx.Abort()
	}
}

func (auth *Authentication) processAuthorization(ctx *gin.Context) *apierrors.APIError {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		return apierrors.NewUnauthorizedError("Unauthorized", "authorization Header not found")
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		return apierrors.NewUnauthorizedError("Unauthorized", "Invalid authorization Header")
	}

	res, err := auth.authServiceClient.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil {
		log.Println("Error requesting to authorization service:", err)
		return apierrors.NewBadGatewayError("Bad Gateway", "Authorization Service not responding")
	}

	if res.Status == http.StatusUnauthorized {
		return apierrors.NewUnauthorizedError("Unauthorized", "Authorization Service response", res.Error)
	}

	if res.Status != http.StatusOK && res.Status != http.StatusCreated {
		return apierrors.NewBadGatewayError("Bad Gateway", "Unknown Authorization Service response")
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()

	return nil
}
