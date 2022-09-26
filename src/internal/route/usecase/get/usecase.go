package get

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Repository interface {
	GetAll() (*[]route.Route, error)
	GetAllEnabled() (*[]route.Route, error)
}

type UseCaseGetRoute struct {
	repository Repository
}

func NewUseCaseGetRoute(repository Repository) *UseCaseGetRoute {
	return &UseCaseGetRoute{repository: repository}
}

func (uc UseCaseGetRoute) Execute(ctx gin.Context) ([]route.Route, error) {
	domainRoute, err := uc.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return *domainRoute, nil
}
