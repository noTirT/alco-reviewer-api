package locations

import "time"

type LocationsService interface {
	CreateLocation(request *CreateLocationRequest) error
    GetLocationsByDrinkId(request *LocationsByDrinkRequest) ([]Location, error)
}

type locationsService struct {
	repo LocationsRepository
}

func NewLocationsService(repo LocationsRepository) LocationsService {
	return &locationsService{
		repo: repo,
	}
}

func (service *locationsService) CreateLocation(request *CreateLocationRequest) error {
	location := &Location{
		Id:        "",
		Name:      request.Name,
		Type:      request.Type,
		Address:   request.Address,
		City:      request.City,
		ZipCode:   request.ZipCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return service.repo.CreateLocation(location)
}

func (service *locationsService) GetLocationsByDrinkId(request *LocationsByDrinkRequest) ([]Location, error) {
    return service.repo.GetLocationsByDrinkId(request.DrinkId)
}
