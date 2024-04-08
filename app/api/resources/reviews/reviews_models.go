package reviews

import "time"

type Review struct {
	Id         string    `json:"id" sql:"id" db:"id"`
	ReviewerId string    `json:"reviewer_id" sql:"reviewer_id" db:"reviewer_id"`
	Rating     int       `json:"rating" sql:"rating" db:"rating"`
	ReviewText string    `json:"review_text" sql:"review_text" db:"review_text"`
	DrinkId    string    `json:"drink_id" sql:"drink_id" db:"drink_id"`
	LocationId string    `json:"location_id" sql:"location_id" db:"location_id"`
	CreatedAt  time.Time `json:"created_at" sql:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" sql:"updated_at" db:"updated_at"`
}
