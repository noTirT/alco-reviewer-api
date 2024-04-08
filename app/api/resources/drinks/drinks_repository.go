package drinks

import (
	"errors"
	"noTirT/alcotracker/app/db"

	"github.com/blockloop/scan/v2"
)

type DrinksRepository interface {
	CreateDrink(drink *Drink, locationId string) error
	GetDrinksByLocationId(locationId string) ([]Drink, error)
	GetDrinkById(drinkId string) (*Drink, error)
}

type drinksRepository struct {
	db *db.PostgresDB
}

func NewDrinksRepository(db *db.PostgresDB) DrinksRepository {
	return &drinksRepository{
		db: db,
	}
}

func (repo *drinksRepository) CreateDrink(drink *Drink, locationId string) error {
	stmt, err := repo.db.Db.Prepare("INSERT INTO drinks(name, alcohol) VALUES($1, $2) RETURNING id;")

	if err != nil {
		return errors.New("Error preparing insert statement")
	}
	defer stmt.Close()

	var drinkId string
	err = stmt.QueryRow(drink.Name, drink.Alcohol).Scan(&drinkId)
	if err != nil {
		return errors.New("Error retrieving drink id")
	}

	_, err = repo.db.Db.Exec("INSERT INTO drinks_to_locations(drink_id, location_id) VALUES($1, $2);", drinkId, locationId)

	return err
}

func (repo *drinksRepository) GetDrinksByLocationId(locationId string) ([]Drink, error) {
	rows, err := repo.db.Db.Query("SELECT d.* FROM drinks d INNER JOIN drinks_to_locations dtl ON d.id = dtl.drink_id WHERE dtl.location_id = $1;", locationId)
	if err != nil {
		return nil, err
	}

	var drinks []Drink
	err = scan.Rows(&drinks, rows)

	return drinks, err
}

func (repo *drinksRepository) GetDrinkById(drinkId string) (*Drink, error) {
	rows, err := repo.db.Db.Query("SELECT * FROM drinks d WHERE id=$1;", drinkId)
	if err != nil {
        return nil, err
	}

	var drink Drink
	err = scan.Row(&drink, rows)

	return &drink, err
}
