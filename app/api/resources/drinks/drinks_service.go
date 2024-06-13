package drinks

import "time"

type DrinksService interface {
	CreateDrink(request *CreateDrinkRequest) error
	GetDrinksByLocationId(locationId string) ([]Drink, error)
}

type drinksService struct {
	repo DrinksRepository
}

func NewDrinksService(repo DrinksRepository) DrinksService {
	return &drinksService{
		repo: repo,
	}
}

func (service *drinksService) CreateDrink(request *CreateDrinkRequest) error {
	drink := &Drink{
		Id:        "",
		Name:      request.Name,
		Alcohol:   request.Alcohol,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return service.repo.CreateDrink(drink, request.LocationId)
}

func (service *drinksService) GetDrinksByLocationId(locationId string) ([]Drink, error) {
	return service.repo.GetDrinksByLocationId(locationId)
}
