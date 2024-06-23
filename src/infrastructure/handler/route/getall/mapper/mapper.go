package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/getall/contract"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Mapper struct {
}

func (m Mapper) ToDTOs(domainRoutes []route.Route) []contract.GetRouteResponse {
	dtoRoutes := make([]contract.GetRouteResponse, 0)
	for _, domainRoute := range domainRoutes {
		dtoMethods := make([]string, 0)
		for _, domainMethod := range domainRoute.Methods() {
			dtoMethods = append(dtoMethods, domainMethod.Value())
		}

		tmpDTO := contract.GetRouteResponse{
			Id:           domainRoute.Id(),
			RelativePath: domainRoute.RelativePath(),
			UrlTarget:    domainRoute.UrlTarget(),
			TypeTarget:   domainRoute.TypeTarget(),
			Secure:       domainRoute.Secure(),
			Enable:       domainRoute.Enable(),
			Methods:      dtoMethods,
		}
		dtoRoutes = append(dtoRoutes, tmpDTO)
	}
	return dtoRoutes
}
