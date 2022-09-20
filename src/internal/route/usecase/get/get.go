package get

import (
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

func (uc UseCaseGetRoute) Execute() (*[]route.Route, error) {
	domainRoute, err := uc.repository.GetAllEnabled()
	if err != nil {
		return nil, err
	}
	return domainRoute, nil
}

func (uc UseCaseGetRoute) Update(routes *[]route.Route) error {
	tmpRoutes, err := uc.repository.GetAllEnabled()
	if err != nil {
		return err
	}
	
	for idx, _ := range *tmpRoutes {
		tmpRouteItem := &(*tmpRoutes)[idx]
		for idy, _ := range *routes {
			routeItem := &(*routes)[idy]
			if tmpRouteItem.Id() == routeItem.Id() {
				routeItem.SetUrlTarget(tmpRouteItem.UrlTarget())
				routeItem.SetEnable(tmpRouteItem.Enable())
			}
		}
	}
	return nil
}
