package auth

import (
	"net/http"
)

func AddRoutes(router *http.ServeMux, controller UserController, middleware AuthHandler) {
    authRouter := http.NewServeMux()

    signUpHandler := http.HandlerFunc(controller.SignUp)
    signInHandler := http.HandlerFunc(controller.SignIn)
    refreshHandler := http.HandlerFunc(controller.RefreshToken)

	authRouter.Handle("POST /signup", middleware.MiddlewareValidateUserInfo(signUpHandler))
	authRouter.Handle("POST /signin", middleware.MiddlewareValidateUserInfo(signInHandler))
    authRouter.Handle("GET /refresh", middleware.MiddlewareValidateRefreshToken(refreshHandler))

    router.Handle("/auth/", http.StripPrefix("/auth", authRouter))
}
