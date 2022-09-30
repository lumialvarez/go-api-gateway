package mapper

import (
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/dao"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Mapper struct{}

func (m Mapper) ToDomain(daoRoute dao.Route) *route.Route {
	return route.NewRoute(daoRoute.Id, daoRoute.RelativePath, daoRoute.UrlTarget, daoRoute.TypeTarget, daoRoute.Secure, daoRoute.Enable)
}
