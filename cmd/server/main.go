package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rumiani/rate-api/internal/db"
	"github.com/rumiani/rate-api/internal/server"
)

func main() {
	_ = godotenv.Load()

	if err := db.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	r := server.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	r.Run(":" + port)
}
