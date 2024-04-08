package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/configs"
	"noTirT/alcotracker/util"
	"strings"
)

type AuthHandler interface {
	MiddlewareValidateUserInfo(next http.Handler) http.Handler
	MiddlewareValidateAccessToken(next http.Handler) http.Handler
	MiddlewareValidateRefreshToken(next http.Handler) http.Handler
}

type authHandler struct {
	configs     *configs.Configuration
	userService AuthService
	validator   *util.Validation
}

func NewAuthHandler(configs *configs.Configuration, userService AuthService) *authHandler {
	return &authHandler{
		configs:     configs,
		userService: userService,
		validator:   util.NewValidation(),
	}
}

func (ah *authHandler) MiddlewareValidateUserInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		UseCORS(w)

		userObj := &shared.User{}

		err := util.FromJSON(userObj, r.Body)
		if err != nil {
			log.Println("deserialization of user json failed", err)
			w.WriteHeader(http.StatusBadRequest)
			util.ToJSON(&shared.GenericResponse{Status: false, Message: err.Error()}, w)
			return
		}

		errs := ah.validator.Validate(userObj)
		if len(errs) != 0 {
			log.Println("Validation of user json failed", errs)
			w.WriteHeader(http.StatusBadRequest)
			util.ToJSON(&shared.GenericResponse{Status: false, Message: strings.Join(errs.Errors(), ",")}, w)
			return
		}

		ctx := context.WithValue(r.Context(), shared.UserKey{}, *userObj)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (ah *authHandler) MiddlewareValidateAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		UseCORS(w)

		token, err := extractToken(r)
		if cancel := shared.HandleTokenParseError(err, w); cancel {
			return
		}

		userID, err := ah.userService.ValidateAccessToken(token)
		if cancel := shared.HandleTokenValidationError(err, w); cancel {
			return
		}

		ctx := context.WithValue(r.Context(), shared.UserIDKey{}, userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (ah *authHandler) MiddlewareValidateRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		UseCORS(w)

		token, err := extractToken(r)
		if cancel := shared.HandleTokenParseError(err, w); cancel {
			return
		}

		userID, customKey, err := ah.userService.ValidateRefreshToken(token)
		if cancel := shared.HandleTokenValidationError(err, w); cancel {
			return
		}

		userObj, err := ah.userService.GetUserByID(userID)
		if err != nil {
			log.Println("invalid token: wrong ID while parsing", err)
			w.WriteHeader(http.StatusBadRequest)
			util.ToJSON(&shared.GenericResponse{Status: false, Message: "Unable to fetch corresponding user"}, w)
			return
		}

		actualCustomKey := util.GenerateCustomKey(fmt.Sprintf("%d", userObj.Id), userObj.TokenHash)
		if customKey != actualCustomKey {
			log.Println("wrong token: authentication failed")
			w.WriteHeader(http.StatusBadRequest)
			util.ToJSON(&shared.GenericResponse{Status: false, Message: "Authentication failed. Invalid token"}, w)
			return
		}

		ctx := context.WithValue(r.Context(), shared.UserKey{}, *userObj)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", errors.New("Token not provided or malformed")
	}
	return authHeaderContent[1], nil
}

func UseCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
