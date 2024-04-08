package locations

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller LocationsController, middleware auth.AuthHandler){
    locationsRouter := http.NewServeMux()

    createLocationHandler := http.HandlerFunc(controller.CreateLocation)
    getLocationsByDrinkHandler := http.HandlerFunc(controller.GetLocationsByDrinkId)

    locationsRouter.Handle("POST /", createLocationHandler)
    locationsRouter.Handle("GET /location-by-drink", getLocationsByDrinkHandler)

    router.Handle("/locations/", http.StripPrefix("/locations", middleware.MiddlewareValidateAccessToken(locationsRouter)))
}
