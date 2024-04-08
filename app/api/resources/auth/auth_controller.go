package auth

import (
	"log"
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"

	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service AuthService
}

func NewAuthController(service AuthService) UserController {
	return &userController{
		service: service,
	}
}

func (controller *userController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(shared.UserKey{}).(shared.User)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("unable to hash password", err)
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&shared.GenericResponse{Status: false, Message: "Unable to create user."}, w)
		return
	}

	user.Password = string(hashedPassword)

	_, err = controller.service.CreateUser(&user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&shared.GenericResponse{Status: false, Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	util.ToJSON(&shared.GenericResponse{Status: true, Message: "shared.User created successfully."}, w)
}

func (controller *userController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userDetails := r.Context().Value(shared.UserKey{}).(shared.User)

	user, err := controller.service.GetUserByUsername(&userDetails)

	if err != nil {
		log.Println(err)
	}

	if valid := controller.service.Authenticate(&userDetails, user); !valid {
		log.Println("Authentication failed")
		w.WriteHeader(http.StatusBadRequest)
		util.ToJSON(&shared.GenericResponse{Status: false, Message: "Incorrect password"}, w)
		return
	}

	accessToken, err := controller.service.GenerateAccessToken(user)

	if err != nil {
		log.Println("Unable to generate access token ", err)
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&shared.GenericResponse{Status: false, Message: "Unable to login the user"}, w)
		return
	}

	refreshToken, err := controller.service.GenerateRefreshToken(user)

	if err != nil {
		log.Println("Unable to generate refresh token", err)
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&shared.GenericResponse{Status: false, Message: "Unable to login the user"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Successfully logged in",
		Data:    &AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken, Username: user.Username},
	}, w)
}

func (controller *userController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := r.Context().Value(shared.UserKey{}).(shared.User)
	accessToken, err := controller.service.GenerateAccessToken(&user)

	if err != nil {
		log.Println("Unable to generate access token", err)
		w.WriteHeader(http.StatusInternalServerError)
		util.ToJSON(&shared.GenericResponse{Status: false, Message: "Unable to generate access token. Please try again later."}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Successfully generated new access token",
		Data:    &AuthResponse{AccessToken: accessToken},
	}, w)
}
