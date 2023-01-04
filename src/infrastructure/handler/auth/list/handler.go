package handlerListAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/list/mapper"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/authentication"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
	"log"
	"net/http"
	"strconv"
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
	id, err := strconv.ParseInt(ctx.Request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		id = 0
	}
	userName := ctx.Request.URL.Query().Get("username")

	serviceRequest := h.mapper.ToServiceRequest(id, userName)

	serviceResponse, err := h.authenticationService.AuthServiceClient.Client.List(ctx, &serviceRequest)
	if err != nil {
		log.Println("Error requesting to authorization service:", err)
		return handlers.ToAPIError(err)
	}
	response := h.mapper.ToDTOResponse(serviceResponse)

	ctx.JSON(http.StatusOK, response)

	return nil
}
