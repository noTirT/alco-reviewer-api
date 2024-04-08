package drinks

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller DrinksController, middleware auth.AuthHandler) {
	drinksRouter := http.NewServeMux()

	createDrinkHandler := http.HandlerFunc(controller.CreateDrink)
	getDrinksByLocationHandler := http.HandlerFunc(controller.GetDrinksByLocationId)

	drinksRouter.Handle("POST /", createDrinkHandler)
	drinksRouter.Handle("GET /drink-by-location", getDrinksByLocationHandler)

	router.Handle("/drinks/", http.StripPrefix("/drinks", middleware.MiddlewareValidateAccessToken(drinksRouter)))
}
