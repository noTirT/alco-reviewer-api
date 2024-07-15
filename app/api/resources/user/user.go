package user

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/db"
)

func UserInit(database *db.PostgresDB, router *http.ServeMux, middleware auth.AuthHandler) {
	repo := NewUserRepository(database)
	service := NewUserService(repo)
	controller := NewUserController(service)

	AddRoutes(router, controller, middleware)
}
