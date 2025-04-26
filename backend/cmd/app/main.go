package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/w0ikid/world-map-tracker/internal/app"
)

func main() {
    // Define config file path
    configFile := flag.String("config", "etc/secrets/.env", "Path to configuration file")
    
    flag.Parse()

    // Try to load .env file, but don't fail if it doesn't exist (for Render deployment)
    _ = godotenv.Load(*configFile)

    log.Println("Config file loaded:", *configFile)

    // Запуск приложения
    if err := app.Run(*configFile); err != nil {
        log.Fatalf("Failed to run application: %v", err)
    }
}