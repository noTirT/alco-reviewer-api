package reviews

import (
	"noTirT/alcotracker/app/db"

	"github.com/blockloop/scan/v2"
)

type ReviewsRepository interface {
	CreateReview(review *Review) error
	GetReviewsByReviewerId(reviewerId string) ([]Review, error)
}

type reviewsRepository struct {
	db *db.PostgresDB
}

func NewReviewsRepository(db *db.PostgresDB) ReviewsRepository {
	return &reviewsRepository{
		db: db,
	}
}

func (repo *reviewsRepository) CreateReview(review *Review) error {
	_, err := repo.db.Db.Exec("INSERT INTO reviews(reviewer_id, rating, review_text, drink_id, location_id) VALUES ($1, $2, $3, $4, $5);", review.ReviewerId, review.Rating, review.ReviewText, review.DrinkId, review.LocationId)

	return err
}

func (repo *reviewsRepository) GetReviewsByReviewerId(reviewerId string) ([]Review, error) {
	rows, err := repo.db.Db.Query("SELECT * FROM reviews WHERE reviewer_id=$1;", reviewerId)

	if err != nil {
		return nil, err
	}

	var reviews []Review
	err = scan.Rows(&reviews, rows)

	return reviews, err
}
