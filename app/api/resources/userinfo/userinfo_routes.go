package userinfo

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller UserinfoController, middleware auth.AuthHandler){
    userinfoRouter := http.NewServeMux()

    changeUsernameHandler := http.HandlerFunc(controller.ChangeUsername)
    getProfileHandler := http.HandlerFunc(controller.GetUserProfile)

    userinfoRouter.Handle("PUT /change-username", changeUsernameHandler)
    userinfoRouter.Handle("GET /profile", getProfileHandler)

    router.Handle("/userinfo/", http.StripPrefix("/userinfo", middleware.MiddlewareValidateAccessToken(userinfoRouter)))
}
