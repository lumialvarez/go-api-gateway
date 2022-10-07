package save

import "github.com/lumialvarez/go-api-gateway/src/internal/route"

type Repository interface {
	Save(route route.Route) error
}

type UseCaseSaveRoute struct {
	repository Repository
}

func NewUseCaseSaveRoute(repository Repository) *UseCaseSaveRoute {
	return &UseCaseSaveRoute{repository: repository}
}

func (uc UseCaseSaveRoute) Execute(route route.Route) error {
	err := uc.repository.Save(route)
	if err != nil {
		return err
	}
	return nil
}
