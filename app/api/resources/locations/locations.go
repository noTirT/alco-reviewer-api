package locations

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/db"
)

func LocationsInit(db *db.PostgresDB, router *http.ServeMux, middleware auth.AuthHandler) {
    repo := NewLocationsRepository(db)
    service := NewLocationsService(repo)
    controller := NewLocationsController(service)

    AddRoutes(router, controller, middleware)
}
