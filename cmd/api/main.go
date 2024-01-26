package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	// db := database.New(conn)
}
