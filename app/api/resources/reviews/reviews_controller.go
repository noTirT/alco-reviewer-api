package reviews

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type ReviewsController interface {
	CreateReview(w http.ResponseWriter, r *http.Request)
	GetReviews(w http.ResponseWriter, r *http.Request)
}

type reviewsController struct {
	service ReviewsService
}

func NewReviewsController(service ReviewsService) ReviewsController {
	return &reviewsController{
		service: service,
	}
}

func (controller *reviewsController) CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createReviewRequest := &CreateReviewRequest{}
	err := util.FromJSON(createReviewRequest, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	userId := r.Context().Value(shared.UserIDKey{}).(string)
	createReviewRequest.ReviewerId = userId

	err = controller.service.CreateReview(createReviewRequest)

	if cancel := shared.HandlePostError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Created review successfully",
	}, w)
}

func (controller *reviewsController) GetReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.Context().Value(shared.UserIDKey{}).(string)
	reviews, err := controller.service.GetReviewsByReviewerId(&GetReviewsRequest{ReviewerId: userId})

	if cancel := shared.HandleGetError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Reviews retrieved successfully",
		Data:    reviews,
	}, w)
}
