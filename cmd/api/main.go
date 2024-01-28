package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("No port found!")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No database url found!")
	}
	_, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	router := chi.NewRouter()
	server := http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}
	log.Printf("Server running on port %v", portStr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
