package locations

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller LocationsController, middleware auth.AuthHandler) {
	createLocationHandler := http.HandlerFunc(controller.CreateLocation)
	getLocationsByDrinkHandler := http.HandlerFunc(controller.GetLocationsByDrinkId)
	getAllLocationsHandler := http.HandlerFunc(controller.GetAllLocations)

	router.Handle("POST /locations", middleware.MiddlewareValidateAccessToken(createLocationHandler))
	router.Handle("GET /locations", middleware.MiddlewareValidateAccessToken(getAllLocationsHandler))
	router.Handle("GET /locations/{drinkId}", middleware.MiddlewareValidateAccessToken(getLocationsByDrinkHandler))
}
