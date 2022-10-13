package mapperGetAllRoutes

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall/contract"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Mapper struct {
}

func (m Mapper) ToDTOs(domainRoutes []route.Route) []contract.GetRouteResponse {
	var dtoRoutes []contract.GetRouteResponse
	for _, domainRoute := range domainRoutes {
		tmpDTO := contract.GetRouteResponse{
			Id:           domainRoute.Id(),
			RelativePath: domainRoute.RelativePath(),
			UrlTarget:    domainRoute.UrlTarget(),
			TypeTarget:   domainRoute.TypeTarget(),
			Secure:       domainRoute.Secure(),
			Enable:       domainRoute.Enable(),
		}
		dtoRoutes = append(dtoRoutes, tmpDTO)
	}
	return dtoRoutes
}
