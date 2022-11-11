package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save/contract"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Mapper struct {
}

func (m Mapper) ToDomain(dtoRoute contract.SaveRouteRequest) route.Route {
	return *route.NewRoute(dtoRoute.Id, dtoRoute.RelativePath, dtoRoute.UrlTarget, dtoRoute.TypeTarget, dtoRoute.Secure, dtoRoute.Enable)
}
