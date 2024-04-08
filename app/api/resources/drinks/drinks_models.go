package drinks

import "time"

type Drink struct {
	Id        string    `json:"id" sql:"id" db:"id"`
	Name      string    `json:"name" sql:"name" db:"name"`
	Alcohol   bool      `json:"alcohol" sql:"alcohol" db:"alcohol"`
	CreatedAt time.Time `json:"created_at" sql:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at" db:"updated_at"`
}
