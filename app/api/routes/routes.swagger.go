package routes

import (
	"github.com/swaggo/http-swagger"
	"net/http"
	_ "noTirT/alcotracker/docs"
)

func SwaggerRoute(router *http.ServeMux) {
	router.Handle("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}
