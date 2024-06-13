package drinks

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller DrinksController, middleware auth.AuthHandler) {
	createDrinkHandler := http.HandlerFunc(controller.CreateDrink)
	getDrinksByLocationHandler := http.HandlerFunc(controller.GetDrinksByLocationId)

	router.Handle("POST /drinks", middleware.MiddlewareValidateAccessToken(createDrinkHandler))
	router.Handle("GET /drinks/{locationId}", middleware.MiddlewareValidateAccessToken(getDrinksByLocationHandler))
}
