package main

import (
	"log"
	"net/http"
	"os"

	"github.com/crnvl96/temperature-calculator/api/handler"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	router := http.NewServeMux()

	handler.NewWeatherHandler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
