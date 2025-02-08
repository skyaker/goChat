package main

import (
	"log"
	"net/http"

	_ "auth_service/docs"
	handlers "auth_service/src/handlers"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	r := chi.NewRouter()

	r.Post("/auth", handlers.CreateToken())

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))

	log.Println("Auth server is running")
	http.ListenAndServe(":8081", r)
}
