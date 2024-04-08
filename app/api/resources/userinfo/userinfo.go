package userinfo

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/db"
)

func UserinfoInit(database *db.PostgresDB, router *http.ServeMux, middleware auth.AuthHandler) {
    repo := NewUserinfoRepository(database)
    service := NewUserinfoService(repo)
    controller := NewUserinfoController(service)

    AddRoutes(router, controller, middleware)
}
