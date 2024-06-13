package auth

import (
	"net/http"
)

func AddRoutes(router *http.ServeMux, controller UserController, middleware AuthHandler) {
	signUpHandler := http.HandlerFunc(controller.SignUp)
	signInHandler := http.HandlerFunc(controller.SignIn)
	refreshHandler := http.HandlerFunc(controller.RefreshToken)

	router.Handle("POST /auth/signup", middleware.MiddlewareValidateUserInfo(signUpHandler))
	router.Handle("POST /auth/signin", middleware.MiddlewareValidateUserInfo(signInHandler))
	router.Handle("GET /auth/refresh", middleware.MiddlewareValidateRefreshToken(refreshHandler))
}
