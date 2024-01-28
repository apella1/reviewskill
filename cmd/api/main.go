package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"reviewskill/config"
	"reviewskill/handlers"
	"reviewskill/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}
	db := database.New(conn)
	apiCfg := config.ApiConfig{
		DB: db,
	}
	handler := handlers.Handler{
		Cfg: &apiCfg,
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://", "http://"},
		AllowedMethods:   []string{"GET", "DELETE", "PUT", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Post("/create_user", handler.HandlerCreateUser)
	router.Mount("/v1", v1Router)
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
