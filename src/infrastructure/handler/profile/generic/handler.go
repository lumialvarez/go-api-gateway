package handlerGenericProfile

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/profile/generic/mapper"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/services/grpc/profile"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"github.com/lumialvarez/go-common-tools/http/handlers"
	"github.com/lumialvarez/go-grpc-profile-service/src/infrastructure/handler/grpc/profile/pb"
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

type ProfileService struct {
	ProfileServiceClient *profile.ServiceClient
}

type Handler struct {
	mapper              mapper.Mapper
	apiResponseProvider ApiResponseProvider
	profileService      ProfileService
}

func NewHandler(apiResponseProvider ApiResponseProvider, profileService ProfileService) *Handler {
	return &Handler{apiResponseProvider: apiResponseProvider, profileService: profileService}
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
	case "list":
		return h.performList(ctx)
	case "save":
		return h.performSave(ctx)
	case "update":
		return h.performUpdate(ctx)
	default:
		return apierrors.NewInternalServerError(internalServerError, "Adequate profile method to perform not found")
	}
}

func (h Handler) performList(ctx *gin.Context) *apierrors.APIError {
	id, err := strconv.ParseInt(ctx.Request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		id = 0
	}
	userName := ctx.Request.URL.Query().Get("language")

	serviceRequest := h.mapper.ToListRequest(id, userName)

	serviceResponse, err := h.profileService.ProfileServiceClient.Client.List(ctx, &serviceRequest)
	return manageResponse(ctx, serviceResponse, err)
}

func (h Handler) performSave(ctx *gin.Context) *apierrors.APIError {
	var request pb.SaveRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceResponse, err := h.profileService.ProfileServiceClient.Client.Save(ctx, &request)
	return manageResponse(ctx, serviceResponse, err)
}

func (h Handler) performUpdate(ctx *gin.Context) *apierrors.APIError {
	var request pb.UpdateRequest
	if err := ctx.BindJSON(&request); err != nil {
		return apierrors.NewBadRequestError(invalidFormat, err.Error())
	}
	serviceResponse, err := h.profileService.ProfileServiceClient.Client.Update(ctx, &request)
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
