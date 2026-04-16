package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"noTirT/alcotracker/app/api/resources/auth"
	"noTirT/alcotracker/app/api/resources/drinks"
	"noTirT/alcotracker/app/api/resources/locations"
	"noTirT/alcotracker/app/api/resources/reviews"
	"noTirT/alcotracker/app/api/resources/user"
	"noTirT/alcotracker/app/api/routes"
	"noTirT/alcotracker/app/db"
	"noTirT/alcotracker/configs"
	"noTirT/alcotracker/docs"
)

//	@title			Alcotracker API
//	@version		1.0
//	@description	API for alcotracker backend
//	@termsOfService	http://swagger.io/terms/

//	@contact.name Tom Manger
//	@contact.email	tommanger55@gmail.com

// @securityDefinitions.apiKey	Bearer
// @in							header
// @name						Authorization
func main() {
	config := configs.NewConfiguration(".env")

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	fmt.Printf("Server starting at port %s\n", config.ServerPort)

	docs.SwaggerInfo.Host = fmt.Sprintf("localhost%s", config.ServerPort)

	router := http.NewServeMux()

	database := db.InitPostgresDB()

	endpoints := []func(*db.PostgresDB, *http.ServeMux, auth.AuthHandler){
		user.UserInit,
		reviews.ReviewsInit,
		drinks.DrinksInit,
		locations.LocationsInit,
	}

	_, middleware := auth.AuthInit(&database, router, config)
	for _, endpoint := range endpoints {
		endpoint(&database, router, middleware)
	}

	routes.SwaggerRoute(router)

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
