package updateRoute

import "github.com/lumialvarez/go-api-gateway/src/internal/route"

type Repository interface {
	Update(route route.Route) error
	GetById(id int64) (*route.Route, error)
}

type UseCaseUpdateRoute struct {
	repository Repository
}

func NewUseCaseUpdateRoute(repository Repository) *UseCaseUpdateRoute {
	return &UseCaseUpdateRoute{repository: repository}
}

func (uc UseCaseUpdateRoute) Execute(route route.Route) error {
	bdRoute, err := uc.repository.GetById(route.Id())
	if err != nil {
		return err
	}

	bdRoute.UpdateRoute(route)

	err = uc.repository.Update(*bdRoute)
	if err != nil {
		return err
	}

	return nil
}
