package route

import (
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

func Init(config config.Config) (Repository, error) {
	postgresqlClient := postgresql.Init(config.DBUrl)
	err := postgresqlClient.DB.AutoMigrate(dao.Route{})
	if err != nil {
		return Repository{}, err
	}
	return Repository{postgresql: postgresqlClient, mapper: mapper.Mapper{}}, nil
}

func (repository *Repository) GetAll() (*[]route.Route, error) {
	var daoRoutes []dao.Route
	result := repository.postgresql.DB.Preload("Methods").Find(&daoRoutes)
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
	result := repository.postgresql.DB.Preload("Methods").Where(&dao.Route{Enable: true}).Find(&daoRoutes)
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
	result := repository.postgresql.DB.Preload("Methods").Where(&dao.Route{Id: id}).First(&daoRoute)
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

	var presentMethods []dao.Method
	result := repository.postgresql.DB.Find(&presentMethods)
	if result.Error != nil {
		return result.Error
	}
	methodsToUpdate := daoRoute.Methods
	for i, inputMethod := range methodsToUpdate {
		isPresent := false
		for _, presentMethod := range presentMethods {
			if inputMethod.Name == presentMethod.Name {
				isPresent = true
				methodsToUpdate[i] = &presentMethod
				break
			}
		}
		if !isPresent {
			methodsToUpdate = append(methodsToUpdate, inputMethod)
		}
	}

	err := repository.postgresql.DB.Model(&daoRoute).Association("Methods").Clear()
	if err != nil {
		return err
	}

	result = repository.postgresql.DB.Preload("Methods").Model(&daoRoute).Where(&dao.Route{Id: daoRoute.Id}).Updates(
		map[string]interface{}{
			"UrlTarget": daoRoute.UrlTarget,
			"Secure":    daoRoute.Secure,
			"Enable":    daoRoute.Enable,
			"Methods":   methodsToUpdate,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
