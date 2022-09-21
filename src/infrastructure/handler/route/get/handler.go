package get

import (
	"github.com/gin-gonic/gin"
)

type Mapper interface {
	//ToDTO(domainClient *cliente.Client) *contract.GetClientResponse
}

type UseCase interface {
	//Execute(ctx context.Context, id string) (*cliente.Client, error)
}

type ApiResponseProvider interface {
	//ToAPIResponse(err error) *apierror.APIError
}

type Handler struct {
	mapper              Mapper
	useCase             UseCase
	apiResponseProvider ApiResponseProvider
}

func (h Handler) Handler(ginCtx *gin.Context) {
	//Controlar los errores
	//handlers.ErrorWrapper(h.handler, ginCtx)
}

func (h Handler) handler(ginCtx *gin.Context) error {
	/*logger := common.Logger(ginCtx)
	idClient := ginCtx.Param("id_client")
	domainClient, err := h.useCase.Execute(ginCtx, idClient)
	if err != nil {

		logger.Error(err.Error(), err, nil...)
		return h.apiResponseProvider.ToAPIResponse(err)

	}
	dtoClient := h.mapper.ToDTO(domainClient)
	ginCtx.JSON(http.StatusOK, dtoClient)*/
	return nil

}
