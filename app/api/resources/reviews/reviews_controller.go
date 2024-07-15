package reviews

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
	"strconv"
)

type ReviewsController interface {
	CreateReview(w http.ResponseWriter, r *http.Request)
	GetReviews(w http.ResponseWriter, r *http.Request)
	GetReviewsSortedWithOffset(w http.ResponseWriter, r *http.Request)
}

type reviewsController struct {
	service ReviewsService
}

func NewReviewsController(service ReviewsService) ReviewsController {
	return &reviewsController{
		service: service,
	}
}

// CreateReview godoc
//	@Id				CreateReview
//	@Summary		Create Review
//	@Description	Create new review
//	@Tags			Review
//	@Produce		json
//	@Param			Review	body		CreateReviewRequest	true	"Review details"
//	@Success		200		{object}	GenericResponse
//	@Failure		500		{object}	GenericResponse	"Server or Database internal error"
//	@Router			/reviews [POST]
//	@Security		Bearer
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

// GetReviews godoc
//	@Id				GetReviews
//	@Summary		Get Reviews
//	@Description	Get all Reviews made by current user
//	@Tags			Review
//	@Produce		json
//	@Success		200	{object}	GenericResponse
//	@Failure		500	{object}	GenericResponse	"Server or Database internal error"
//	@Router			/reviews/user [GET]
//	@Security		Bearer
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

// GetReviewsSortedWithOffset godoc
//	@Id				GetReviewsSortedWithOffset
//	@Summary		Get Reviews Sorted With Offset
//	@Description	Get all Reviews made by friends of the current user sorted by time and witha given offset for pagination
//	@Tags			Review
//	@Produce		json
//	@Param			offset	query		int	true	"Current offset representing the page"
//	@Param			count	query		int	true	"How many reviews are included in one page"
//	@Success		200		{object}	GenericResponse
//	@Failure		500		{object}	GenericResponse	"Server or Database internal error"
//	@Router			/reviews/feed [GET]
//	@Security		Bearer
func (controller *reviewsController) GetReviewsSortedWithOffset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if shared.HandleParseError(err, w) {
		return
	}

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if shared.HandleParseError(err, w) {
		return
	}

	request := &GetReviewsSortedOffsetRequest{
		Count:  count,
		Offset: offset,
	}

	userId := r.Context().Value(shared.UserIDKey{}).(string)
	reviews, err := controller.service.GetReviewsSortedWithOffset(request, userId)

	if shared.HandleGetError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Reviews retrieved successfully",
		Data:    reviews,
	}, w)
}
