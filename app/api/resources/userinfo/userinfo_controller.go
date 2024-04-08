package userinfo

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type UserinfoController interface {
	ChangeUsername(w http.ResponseWriter, r *http.Request)
	GetUserProfile(w http.ResponseWriter, r *http.Request)
}

type userinfoController struct {
	service UserinfoService
}

func NewUserinfoController(service UserinfoService) UserinfoController {
	return &userinfoController{
		service: service,
	}
}

func (controller *userinfoController) ChangeUsername(w http.ResponseWriter, r *http.Request) {
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

func (controller *userinfoController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := r.Context().Value(shared.UserIDKey{}).(string)
	currUser, err := controller.service.GetUserByID(userID)

	if cancel := shared.HandleGetError(err, w); cancel {
		return
	}
	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Successfully retrieve user profile",
		Data:    currUser,
	}, w)
}
