package drinks

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/db"
)

func DrinksInit(db *db.PostgresDB, router *http.ServeMux, middleware auth.AuthHandler) {
    repo := NewDrinksRepository(db)
    service := NewDrinksService(repo)
    controller := NewDrinksController(service)

    AddRoutes(router, controller, middleware)
}
