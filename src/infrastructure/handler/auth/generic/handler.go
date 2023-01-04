package handlerGenericAuth

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/auth/generic/mapper"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/authentication"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
	"github.com/lumialvarez/go-grpc-auth-service/src/infrastructure/handler/grpc/auth/pb"
	"log"
	"net/http"
	"strconv"
)

const (
	invalidFormat       string = "Invalid message format"
	internalServerError        = "Internal Server Error"
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

func (h Handler) Handler(ginCtx *gin.Context, params ...string) {
	handlers.ErrorWrapperParams(h.handler, ginCtx, params...)
}

func (h Handler) handler(ctx *gin.Context, params ...string) *apierrors.APIError {
	if len(params) != 1 {
		return apierrors.NewInternalServerError(internalServerError, "Wrong params in Generic Authorization handler")
	}

	methodNameToPerform := params[0]

	switch methodNameToPerform {
	case "login":
		return h.performLogin(ctx)
	case "register":
		return h.performRegister(ctx)
	case "update":
		return h.performUpdate(ctx)
	case "validate":
		return h.performValidate(ctx)
	case "list":
		return h.performList(ctx)
	default:
		return apierrors.NewInternalServerError(internalServerError, "Adequate authorization method to perform not found")
	}
}

func (h Handler) performLogin(ctx *gin.Context) *apierrors.APIError {
	var request pb.LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceResponse, err := h.authenticationService.AuthServiceClient.Client.Login(ctx, &request)
	return manageResponse(ctx, serviceResponse, err)
}

func (h Handler) performRegister(ctx *gin.Context) *apierrors.APIError {
	var request pb.RegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceResponse, err := h.authenticationService.AuthServiceClient.Client.Register(ctx, &request)
	return manageResponse(ctx, serviceResponse, err)
}

func (h Handler) performUpdate(ctx *gin.Context) *apierrors.APIError {
	var request pb.UpdateRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceResponse, err := h.authenticationService.AuthServiceClient.Client.Update(ctx, &request)
	return manageResponse(ctx, serviceResponse, err)
}

func (h Handler) performValidate(ctx *gin.Context) *apierrors.APIError {
	var request pb.ValidateRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceResponse, err := h.authenticationService.AuthServiceClient.Client.Validate(ctx, &request)
	return manageResponse(ctx, serviceResponse, err)
}

func (h Handler) performList(ctx *gin.Context) *apierrors.APIError {
	id, err := strconv.ParseInt(ctx.Request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		id = 0
	}
	userName := ctx.Request.URL.Query().Get("username")

	serviceRequest := h.mapper.ToListRequest(id, userName)

	serviceResponse, err := h.authenticationService.AuthServiceClient.Client.List(ctx, &serviceRequest)
	return manageResponse(ctx, serviceResponse, err)
}

func manageResponse(ctx *gin.Context, response interface{}, err error) *apierrors.APIError {
	if err != nil {
		log.Println("Error requesting to authorization service:", err)
		return handlers.ToAPIError(err)
	}
	ctx.JSON(http.StatusOK, response)
	return nil
}
