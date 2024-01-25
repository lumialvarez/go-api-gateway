package route

import (
	"fmt"
	"github.com/lumialvarez/go-api-gateway/src/cmd/devapi/config"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/dao"
	"github.com/lumialvarez/go-api-gateway/src/infrastructure/repository/postgresql/route/mapper"
	"github.com/lumialvarez/go-api-gateway/src/internal/route"
	"github.com/lumialvarez/go-common-tools/platform/postgresql"
)

type Repository struct {
	postgresql postgresql.Client
	mapper     mapper.Mapper
}

func Init(config config.Config) Repository {
	urlDataConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.DBUser, config.DBPassword, config.DBUrl, config.DBPort, config.DBName)
	postgresqlClient := postgresql.Init(urlDataConnection)
	postgresqlClient.DB.AutoMigrate(dao.Route{})
	return Repository{postgresql: postgresqlClient, mapper: mapper.Mapper{}}
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

func (repository *Repository) GetById(id int64) (*route.Route, error) {
	var daoRoute dao.Route
	result := repository.postgresql.DB.Where(&dao.Route{Id: id}).First(&daoRoute)
	if result.Error != nil {
		return nil, result.Error
	}
	domainRoute := repository.mapper.ToDomain(daoRoute)
	return domainRoute, nil
}

func (repository *Repository) Save(route route.Route) error {
	daoRoute := repository.mapper.ToDao(route)
	daoRoute.Id = 0
	result := repository.postgresql.DB.Create(&daoRoute)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *Repository) Update(route route.Route) error {
	daoRoute := repository.mapper.ToDao(route)

	result := repository.postgresql.DB.Model(&daoRoute).Where(&dao.Route{Id: daoRoute.Id}).Updates(
		map[string]interface{}{
			"UrlTarget": daoRoute.UrlTarget,
			"Secure":    daoRoute.Secure,
			"Enable":    daoRoute.Enable,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
