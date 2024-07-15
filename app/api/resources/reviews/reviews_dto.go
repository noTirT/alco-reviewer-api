package reviews

import (
	"noTirT/alcotracker/app/api/resources/drinks"
	"noTirT/alcotracker/app/api/resources/locations"
	"time"
)

type CreateReviewRequest struct {
	ReviewerId string `json:"reviewer_id"`
	Rating     int    `json:"rating" validate:"required"`
	ReviewText string `json:"review_text" validate:"required"`
	DrinkId    string `json:"drink_id"`
	LocationId string `json:"location_id"`
}

type GetReviewsRequest struct {
	ReviewerId string `json:"reviewer_id"`
}

type ResolvedReviewResponse struct {
	ReviewId   string             `json:"review_id"`
	ReviewerId string             `json:"reviewer_id"`
	Rating     int                `json:"rating"`
	ReviewText string             `json:"review_text"`
	Drink      drinks.Drink       `json:"drink"`
	Location   locations.Location `json:"location"`
	CreatedAt  time.Time          `json:"created_at"`
}

type GetReviewsSortedOffsetRequest struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type FeedReviewResponse struct {
	ReviewId     string             `json:"review_id"`
	ReviewerId   string             `json:"reviewer_id"`
	Rating       int                `json:"rating"`
	ReviewText   string             `json:"review_text"`
	ReviewerName string             `json:"reviewer_name"`
	CreatedAt    time.Time          `json:"created_at"`
	Drink        drinks.Drink       `json:"drink"`
	Location     locations.Location `json:"location"`
}
