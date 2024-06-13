package userinfo

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller UserinfoController, middleware auth.AuthHandler) {
	changeUsernameHandler := http.HandlerFunc(controller.ChangeUsername)
	getProfileHandler := http.HandlerFunc(controller.GetUserProfile)

	router.Handle("PUT /userinfo/username", middleware.MiddlewareValidateAccessToken(changeUsernameHandler))
	router.Handle("GET /userinfo", middleware.MiddlewareValidateAccessToken(getProfileHandler))
}
