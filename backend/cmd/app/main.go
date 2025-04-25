package main

import (
	"flag"
	"log"
	"github.com/joho/godotenv"
	"github.com/w0ikid/world-map-tracker/internal/app"
)

func main() {
	// Define config file path
	configFile := flag.String("config", "configs/.env", "Path to configuration file")
	flag.Parse()

	// Load .env file
	err := godotenv.Load(*configFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if err := app.Run(*configFile); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}