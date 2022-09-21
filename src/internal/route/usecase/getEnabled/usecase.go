package getEnabled

import (
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Repository interface {
	GetAll() (*[]route.Route, error)
	GetAllEnabled() (*[]route.Route, error)
}

type UseCaseGetEnabledRoute struct {
	repository Repository
}

func NewUseCaseGetEnabledRoute(repository Repository) *UseCaseGetEnabledRoute {
	return &UseCaseGetEnabledRoute{repository: repository}
}

func (uc UseCaseGetEnabledRoute) Execute() (*[]route.Route, error) {
	domainRoute, err := uc.repository.GetAllEnabled()
	if err != nil {
		return nil, err
	}
	return domainRoute, nil
}

func (uc UseCaseGetEnabledRoute) Update(routes *[]route.Route) error {
	tmpRoutes, err := uc.repository.GetAll()
	if err != nil {
		return err
	}

	for idx := range *tmpRoutes {
		tmpRouteItem := &(*tmpRoutes)[idx]
		for idy := range *routes {
			routeItem := &(*routes)[idy]
			if tmpRouteItem.Id() == routeItem.Id() {
				routeItem.SetUrlTarget(tmpRouteItem.UrlTarget())
				routeItem.SetEnable(tmpRouteItem.Enable())
			}
		}
	}
	return nil
}
