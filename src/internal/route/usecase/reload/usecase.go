package reloadRoute

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

func NewUseCaseReloadRoute(repository Repository) *UseCaseGetRoute {
	return &UseCaseGetRoute{repository: repository}
}

func (uc UseCaseGetRoute) Execute(ctx gin.Context, routes *[]route.Route) ([]route.Route, error) {
	tmpRoutes, err := uc.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for idx := range *tmpRoutes {
		tmpRouteItem := &(*tmpRoutes)[idx]
		for _, routeItem := range *routes {
			if tmpRouteItem.Id() == routeItem.Id() {
				routeItem.UpdateRoute(*tmpRouteItem)
			}
		}
	}
	return *routes, nil
}
