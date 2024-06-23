package mapper

import (
	"errors"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/handler/route/save/contract"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-api-gateway/src/internal/route/enum"
)

const (
	invalidMethodName = "invalid method name"
)

type Mapper struct {
}

func (m Mapper) ToDomain(dtoRoute contract.SaveRouteRequest) (route.Route, error) {
	domainMethods := make([]enum.Method, 0)
	for _, dtoMethod := range dtoRoute.Methods {
		domainMethod := enum.Method(dtoMethod)
		if !domainMethod.IsValid() {
			return route.Route{}, errors.New(invalidMethodName)
		}
		domainMethods = append(domainMethods, domainMethod)
	}

	return *route.NewRoute(dtoRoute.Id, dtoRoute.RelativePath, dtoRoute.UrlTarget, dtoRoute.TypeTarget, dtoRoute.Secure, dtoRoute.Enable, domainMethods), nil
}
