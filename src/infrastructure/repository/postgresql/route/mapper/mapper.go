package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/dao"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-api-gateway/src/internal/route/enum"
)

type Mapper struct{}

func (m Mapper) ToDomain(daoRoute dao.Route) *route.Route {
	methods := make([]enum.Method, 0)
	for _, daoMethod := range daoRoute.Methods {
		methods = append(methods, enum.Method(daoMethod.Name))
	}
	return route.NewRoute(daoRoute.Id, daoRoute.RelativePath, daoRoute.UrlTarget, daoRoute.TypeTarget, daoRoute.Secure, daoRoute.Enable, methods)
}

func (m Mapper) ToDao(domainRoute route.Route) dao.Route {
	daoMethods := make([]*dao.Method, 0)
	for _, domainMethod := range domainRoute.Methods() {
		daoMethods = append(daoMethods, &dao.Method{Name: domainMethod.Value()})
	}

	return dao.Route{
		Id:           domainRoute.Id(),
		RelativePath: domainRoute.RelativePath(),
		UrlTarget:    domainRoute.UrlTarget(),
		TypeTarget:   domainRoute.TypeTarget(),
		Secure:       domainRoute.Secure(),
		Enable:       domainRoute.Enable(),
		Methods:      daoMethods,
	}
}
