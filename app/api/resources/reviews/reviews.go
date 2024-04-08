package reviews

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/api/resources/drinks"
	"noTirT/alcotracker/app/api/resources/locations"
	"noTirT/alcotracker/app/db"
)

func ReviewsInit(db *db.PostgresDB, router *http.ServeMux, middleware auth.AuthHandler) {
	repo := NewReviewsRepository(db)
	drinkRepo := drinks.NewDrinksRepository(db)
	locationRepo := locations.NewLocationsRepository(db)
	service := NewReviewsService(repo, locationRepo, drinkRepo)
	controller := NewReviewsController(service)

	AddRoutes(router, controller, middleware)
}
