package reviews

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller ReviewsController, middleware auth.AuthHandler) {
    reviewsRouter := http.NewServeMux()

    createReviewHandler := http.HandlerFunc(controller.CreateReview)
    getReviewsHandler := http.HandlerFunc(controller.GetReviews)

    reviewsRouter.Handle("POST /", createReviewHandler)
    reviewsRouter.Handle("GET /user", getReviewsHandler)

    router.Handle("/reviews/", http.StripPrefix("/reviews", middleware.MiddlewareValidateAccessToken(reviewsRouter)))
}
