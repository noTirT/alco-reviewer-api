package reviews

import (
	"noTirT/alcotracker/app/api/resources/drinks"
	"noTirT/alcotracker/app/api/resources/locations"
	"time"
)

type ReviewsService interface {
	CreateReview(request *CreateReviewRequest) error
	GetReviewsByReviewerId(request *GetReviewsRequest) ([]ResolvedReviewResponse, error)
}

type reviewsService struct {
	repo         ReviewsRepository
	locationRepo locations.LocationsRepository
	drinkRepo    drinks.DrinksRepository
}

func NewReviewsService(repo ReviewsRepository, locationRepo locations.LocationsRepository, drinkRepo drinks.DrinksRepository) ReviewsService {
	return &reviewsService{
		repo:         repo,
		locationRepo: locationRepo,
		drinkRepo:    drinkRepo,
	}
}

func (service *reviewsService) CreateReview(request *CreateReviewRequest) error {
	review := &Review{
		Id:         "",
		ReviewerId: request.ReviewerId,
		Rating:     request.Rating,
		ReviewText: request.ReviewText,
		DrinkId:    request.DrinkId,
		LocationId: request.LocationId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return service.repo.CreateReview(review)
}

func (service *reviewsService) GetReviewsByReviewerId(request *GetReviewsRequest) ([]ResolvedReviewResponse, error) {
	reviews, err := service.repo.GetReviewsByReviewerId(request.ReviewerId)

	var resultSet []ResolvedReviewResponse

	for _, review := range reviews {
		location, err := service.locationRepo.GetLocationById(review.LocationId)
		if err != nil {
			return nil, err
		}
		drink, err := service.drinkRepo.GetDrinkById(review.DrinkId)
		if err != nil {
			return nil, err
		}

		resolvedReview := ResolvedReviewResponse{
			ReviewId:   review.Id,
			ReviewerId: review.ReviewerId,
			Rating:     review.Rating,
			ReviewText: review.ReviewText,
			Drink:      *drink,
			Location:   *location,
			CreatedAt:  review.CreatedAt,
		}
		resultSet = append(resultSet, resolvedReview)
	}

	return resultSet, err
}
