package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	dbconn "users_service/database/db_connection"
	handlers "users_service/src/handlers"

	"github.com/go-chi/chi/v5"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	_ "users_service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	var db *sql.DB = dbconn.GetDbConnection()
	defer db.Close()

	r := chi.NewRouter()

	r.Post("/register", handlers.AddUser(db))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	log.Println("Server is running")
	http.ListenAndServe(":8080", r)
}
