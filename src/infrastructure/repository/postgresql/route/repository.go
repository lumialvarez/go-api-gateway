package route

import (
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/platform/postgresql"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/dao"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/mapper"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
)

type Repository struct {
	postgresql postgresql.Client
	mapper     mapper.Mapper
}

func Init(config config.Config) Repository {
	return Repository{postgresql: postgresql.Init(config.DBUrl), mapper: mapper.Mapper{}}
}

func (repository *Repository) GetAll() (*[]route.Route, error) {
	var daoRoutes []dao.Route
	result := repository.postgresql.DB.Find(&daoRoutes)
	if result.Error != nil {
		return nil, result.Error
	}
	domainRoutes := make([]route.Route, 0)
	for _, daoRoute := range daoRoutes {
		domainRoutes = append(domainRoutes, *repository.mapper.ToDomain(daoRoute))
	}
	return &domainRoutes, nil
}

func (repository *Repository) GetAllEnabled() (*[]route.Route, error) {

	var daoRoutes []dao.Route
	result := repository.postgresql.DB.Where(&dao.Route{Enable: true}).Find(&daoRoutes)
	if result.Error != nil {
		return nil, result.Error
	}
	domainRoutes := make([]route.Route, 0)
	for _, daoRoute := range daoRoutes {
		domainRoutes = append(domainRoutes, *repository.mapper.ToDomain(daoRoute))
	}
	return &domainRoutes, nil
}
