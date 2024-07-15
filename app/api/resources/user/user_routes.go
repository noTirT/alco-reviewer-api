package user

import (
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
)

func AddRoutes(router *http.ServeMux, controller UserController, middleware auth.AuthHandler) {
	changeUsernameHandler := http.HandlerFunc(controller.ChangeUsername)
	followUserHandler := http.HandlerFunc(controller.FollowUser)
	unfollowUserHandler := http.HandlerFunc(controller.UnfollowUser)
	getUserProfileHandler := http.HandlerFunc(controller.GetUserProfile)

	router.Handle("PUT /user/username", middleware.MiddlewareValidateAccessToken(changeUsernameHandler))
	router.Handle("POST /user/follow/{userId}", middleware.MiddlewareValidateAccessToken(followUserHandler))
	router.Handle("POST /user/unfollow/{userId}", middleware.MiddlewareValidateAccessToken(unfollowUserHandler))
	router.Handle("GET /user/{userId...}", middleware.MiddlewareValidateAccessToken(getUserProfileHandler))
}
