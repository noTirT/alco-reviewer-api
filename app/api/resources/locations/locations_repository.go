package locations

import (
	"github.com/blockloop/scan/v2"
	"noTirT/alcotracker/app/db"
)

type LocationsRepository interface {
	CreateLocation(location *Location) error
	GetLocationsByDrinkId(drinkId string) ([]Location, error)
	GetLocationById(locationId string) (*Location, error)
	GetAllLocations() ([]Location, error)
}

type locationsRepository struct {
	db *db.PostgresDB
}

func NewLocationsRepository(db *db.PostgresDB) LocationsRepository {
	return &locationsRepository{
		db: db,
	}
}

func (repo *locationsRepository) GetAllLocations() ([]Location, error) {
	rows, err := repo.db.Db.Query("SELECT * FROM locations")
	if err != nil {
		return nil, err
	}

	var locations []Location
	err = scan.Rows(&locations, rows)

	return locations, err
}

func (repo *locationsRepository) CreateLocation(location *Location) error {
	_, err := repo.db.Db.Exec("insert into locations(name, type, address, city, zip_code) values($1, $2, $3, $4, $5);", location.Name, location.Type, location.Address, location.City, location.ZipCode)

	return err
}

func (repo *locationsRepository) GetLocationsByDrinkId(drinkId string) ([]Location, error) {
	rows, err := repo.db.Db.Query("SELECT l.* from locations l INNER JOIN drinks_to_locations dtl ON l.id=dtl.location_id WHERE dtl.drink_id=$1", drinkId)
	if err != nil {
		return nil, err
	}

	var locations []Location
	err = scan.Rows(&locations, rows)

	return locations, err
}

func (repo *locationsRepository) GetLocationById(locationId string) (*Location, error) {
	rows, err := repo.db.Db.Query("SELECT * FROM locations WHERE id=$1;", locationId)
	if err != nil {
		return nil, err
	}

	var location Location
	err = scan.Row(&location, rows)

	return &location, err
}
