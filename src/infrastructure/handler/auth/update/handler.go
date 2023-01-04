package handlerUpdateAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/update/contract"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/update/mapper"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/authentication"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
	"log"
	"net/http"
)

const (
	invalidFormat string = "invalid_message_format"
)

type ApiResponseProvider interface {
	ToAPIResponse(err error, cause ...string) *apierrors.APIError
}

type Authentication struct {
	AuthServiceClient *authentication.ServiceClient
}

type Handler struct {
	mapper                mapper.Mapper
	apiResponseProvider   ApiResponseProvider
	authenticationService Authentication
}

func NewHandler(apiResponseProvider ApiResponseProvider, authenticationService Authentication) *Handler {
	return &Handler{apiResponseProvider: apiResponseProvider, authenticationService: authenticationService}
}

func (h Handler) Handler(ginCtx *gin.Context) {
	handlers.ErrorWrapper(h.handler, ginCtx)
}

func (h Handler) handler(ctx *gin.Context) *apierrors.APIError {
	var request contract.UpdateAuthRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceRequest := h.mapper.ToServiceRequest(request)

	_, err := h.authenticationService.AuthServiceClient.Client.Update(ctx, &serviceRequest)
	if err != nil {
		log.Println("Error requesting to authorization service:", err)
		return handlers.ToAPIError(err)
	}

	ctx.Status(http.StatusOK)

	return nil
}
