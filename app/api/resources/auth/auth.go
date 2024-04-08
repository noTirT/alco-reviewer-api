package auth

import (
	"net/http"
	"noTirT/alcotracker/app/db"
	"noTirT/alcotracker/configs"
)

func AuthInit(database *db.PostgresDB, router *http.ServeMux, config *configs.Configuration) (AuthService, AuthHandler) {
	repo := NewAuthRepository(database)
	service := NewAuthService(repo, config)
	controller := NewAuthController(service)
	authMiddleware := NewAuthHandler(config, service)

	AddRoutes(router, controller, authMiddleware)

	return service, authMiddleware
}
