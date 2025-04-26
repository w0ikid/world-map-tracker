package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/w0ikid/world-map-tracker/internal/app"
)

func main() {

    name := os.Getenv("DBnamedb")
    
    db :=os.Getenv("USERdb")

    fmt.Println("Starting application...")
    fmt.Println("DB name:", name)
    fmt.Println("DB user:", db)

    // Define config file path
    configFile := flag.String("config", "configs/.env", "Path to configuration file")
    flag.Parse()

    // Try to load .env file, but don't fail if it doesn't exist (for Render deployment)
    _ = godotenv.Load(*configFile)
    
    // Запуск приложения
    if err := app.Run(*configFile); err != nil {
        log.Fatalf("Failed to run application: %v", err)
    }
}