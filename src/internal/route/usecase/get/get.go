package get

import (
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Repository interface {
	GetAll() ([]route.Route, error)
	GetAllEnabled() ([]route.Route, error)
}

type UseCaseGetRoute struct {
	repository Repository
}

func NewUseCaseGetRoute(repository Repository) *UseCaseGetRoute {
	return &UseCaseGetRoute{repository: repository}
}

func (uc UseCaseGetRoute) Execute() ([]route.Route, error) {
	domainRoute, err := uc.repository.GetAllEnabled()
	if err != nil {
		return nil, err
	}
	return domainRoute, nil
}
