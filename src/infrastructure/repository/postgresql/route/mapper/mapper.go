package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/dao"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Mapper struct{}

func (m Mapper) ToDomain(daoRoute dao.Route) *route.Route {
	return route.NewRoute(daoRoute.Id, daoRoute.RelativePath, daoRoute.UrlTarget, daoRoute.TypeTarget, daoRoute.Secure, daoRoute.Enable)
}

func (m Mapper) ToDao(domainRoute route.Route) dao.Route {
	return dao.Route{
		Id:           domainRoute.Id(),
		RelativePath: domainRoute.RelativePath(),
		UrlTarget:    domainRoute.UrlTarget(),
		TypeTarget:   domainRoute.TypeTarget(),
		Secure:       domainRoute.Secure(),
		Enable:       domainRoute.Enable(),
	}
}
