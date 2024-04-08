package shared

import (
	"log"
	"net/http"
	"noTirT/alcotracker/util"
)

func HandleParseError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Println("Error parsing JSON", err)
	w.WriteHeader(http.StatusBadRequest)
	util.ToJSON(&GenericResponse{Status: false, Message: "Failed to parse request body. Try again later"}, w)
	return true
}

func HandleGetError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Println("Error retrieving data", err)
	w.WriteHeader(http.StatusInternalServerError)
	util.ToJSON(&GenericResponse{Status: false, Message: "Failed to retrieve data. Try again later"}, w)
	return true
}

func HandlePostError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Println("Error retrieving data", err)
	w.WriteHeader(http.StatusInternalServerError)
	util.ToJSON(&GenericResponse{Status: false, Message: "Failed to retrieve data. Try again later"}, w)
	return true
}

func HandleTokenParseError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Println("Token not provided or malformed")
	w.WriteHeader(http.StatusBadRequest)
	util.ToJSON(&GenericResponse{Status: false, Message: "Authentication failed. Invalid token"}, w)
	return true
}

func HandleTokenValidationError(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	log.Println("token validation failed ", "error", err)
	w.WriteHeader(401)
	util.ToJSON(&GenericResponse{Status: false, Message: "Token expired"}, w)
	return true
}
