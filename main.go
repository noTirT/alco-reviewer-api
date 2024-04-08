package main

import (
	"fmt"
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/api/resources/drinks"
	"noTirT/alcotracker/app/api/resources/locations"
	"noTirT/alcotracker/app/api/resources/reviews"
	"noTirT/alcotracker/app/api/resources/userinfo"
	"noTirT/alcotracker/app/db"
	"noTirT/alcotracker/configs"

	"github.com/rs/cors"
	//"noTirT/alcotracker/util"
)

func main() {
	//util.GenerateRSAKeyPairs("auth")
	//util.GenerateRSAKeyPairs("refresh")

	// reset tokenhash for users in a regular fashion
	config := configs.NewConfiguration()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	fmt.Printf("Server starting at port %s\n", config.ServerPort)

	router := http.NewServeMux()

	database := db.InitPostgresDB()

	endpoints := []func(*db.PostgresDB, *http.ServeMux, auth.AuthHandler){
		userinfo.UserinfoInit,
		reviews.ReviewsInit,
		drinks.DrinksInit,
		locations.LocationsInit,
	}

	_, middleware := auth.AuthInit(&database, router, config)
	for _, endpoint := range endpoints {
		endpoint(&database, router, middleware)
	}

	server := http.Server{
		Addr:    config.ServerPort,
		Handler: cors.Handler(corsMiddle(router)),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func corsMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
		next.ServeHTTP(w, r)
	})
}
