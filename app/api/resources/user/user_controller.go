package user

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type UserController interface {
	ChangeUsername(w http.ResponseWriter, r *http.Request)
	GetUserProfile(w http.ResponseWriter, r *http.Request)
	FollowUser(w http.ResponseWriter, r *http.Request)
	UnfollowUser(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &userController{
		service: service,
	}
}

// GetUserById godoc
//
//	@Id				GetUserById
//	@Summary		Get User By Id
//	@Description	Get the User with the given ID. If the ID is left empty, the currently logged in user is retrieved instead
//	@Tags			Userinfo
//	@Produce		json
//	@Param			userId	path		string	true	"ID of the user to retrieve"
//	@Success		200		{object}	GenericResponse
//	@Failure		500		{object}	GenericResponse	"Server or Database internal error"
//	@Router			/user/{userId} [GET]
//	@Security		Bearer
func (controller *userController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.PathValue("userId")
	ownUserId := r.Context().Value(shared.UserIDKey{}).(string)

	var currUser UserProfileFollowingResponse
	var err error

	if userId != "" {
		currUser, err = controller.service.GetUserByID(userId, ownUserId)
	} else {
		currUser, err = controller.service.GetUserByID(ownUserId, ownUserId)
	}

	if shared.HandleGetError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "User successfully retrieved",
		Data:    currUser,
	}, w)
}

// UnfollowUser godoc
//
//	@Id				UnfollowUser
//	@Summary		Unfollow User
//	@Description	Unfollow a user
//	@Tags			Userinfo
//	@Produce		json
//	@Param			userId	path		string	true	"ID of user to unfollow"
//	@Success		200		{object}	GenericResponse
//	@Failure		500		{object}	GenericResponse	"Server or Database internal error"
//	@Router			/user/unfollow/{userId} [POST]
//	@Security		Bearer
func (controller *userController) UnfollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	unfollowId := r.PathValue("userId")
	userID := r.Context().Value(shared.UserIDKey{}).(string)

	err := controller.service.UnfollowUser(userID, unfollowId)

	if shared.HandlePostError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Successfully unfollowed user",
	}, w)
}

// FollowUser godoc
//
//	@Id				FollowUser
//	@Summary		Follow User
//	@Description	Follow a user
//	@Tags			Userinfo
//	@Produce		json
//	@Param			userId	path		string	true	"ID of user to follow"
//	@Success		200		{object}	GenericResponse
//	@Failure		500		{object}	GenericResponse	"Server or Database internal error"
//	@Router			/user/follow/{userId} [POST]
//	@Security		Bearer
func (controller *userController) FollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	followId := r.PathValue("userId")
	userID := r.Context().Value(shared.UserIDKey{}).(string)

	err := controller.service.FollowUser(userID, followId)

	if shared.HandlePostError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Successfully followed user",
	}, w)
}

// ChangeUsername godoc
//
//	@Id				ChangeUsernamce
//	@Summary		Change Username
//	@Description	Update the username of current user
//	@Tags			Userinfo
//	@Produce		json
//	@Param			UsernameChangeRequest	body		UsernameChangeRequest	true	"New username info"
//	@Success		200						{object}	GenericResponse
//	@Failure		500						{object}	GenericResponse	"Server or Database internal error"
//	@Router			/user/change-username [PUT]
//	@Security		Bearer
func (controller *userController) ChangeUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usernameChangeRequest := &UsernameChangeRequest{}
	err := util.FromJSON(usernameChangeRequest, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	userID := r.Context().Value(shared.UserIDKey{}).(string)

	err = controller.service.ChangeUsername(userID, usernameChangeRequest)

	if cancel := shared.HandlePostError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Successfully updated username",
	}, w)
}
