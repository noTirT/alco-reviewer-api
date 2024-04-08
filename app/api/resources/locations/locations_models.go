package locations

import "time"

type Location struct {
	Id        string    `json:"id" sql:"id" db:"id"`
	Name      string    `json:"name" sql:"name" db:"name"`
	Type      string    `json:"type" sql:"type" db:"type"`
	Address   string    `json:"address" sql:"address" db:"address"`
	City      string    `json:"city" sql:"city" db:"city"`
	ZipCode   string    `json:"zip_code" sql:"zip_code" db:"zip_code"`
	CreatedAt time.Time `json:"created_at" sql:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" sql:"updated_at" db:"updated_at"`
}
