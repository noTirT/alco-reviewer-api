package reviews

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller ReviewsController, middleware auth.AuthHandler) {
	createReviewHandler := http.HandlerFunc(controller.CreateReview)
	getReviewsHandler := http.HandlerFunc(controller.GetReviews)
	getReviewsSortedWithOffset := http.HandlerFunc(controller.GetReviewsSortedWithOffset)

	router.Handle("POST /reviews", middleware.MiddlewareValidateAccessToken(createReviewHandler))
	router.Handle("GET /reviews/user", middleware.MiddlewareValidateAccessToken(getReviewsHandler))
	router.Handle("GET /reviews/feed", middleware.MiddlewareValidateAccessToken(getReviewsSortedWithOffset))
}
