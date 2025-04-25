package main

import (
	"flag"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/w0ikid/world-map-tracker/internal/app"
)

func main() {
	// Define config file path
	configFile := flag.String("config", "configs/.env", "Path to configuration file")
	flag.Parse()

	// Try to load .env only if file exists
	if _, err := os.Stat(*configFile); err == nil {
		if err := godotenv.Load(*configFile); err != nil {
			log.Printf("Warning: error loading .env file: %v", err)
		} else {
			log.Println(".env file loaded")
		}
	} else {
		log.Println("No .env file found, using environment variables")
	}

	if err := app.Run(*configFile); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}