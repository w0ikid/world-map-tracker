package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/w0ikid/world-map-tracker/internal/app"
)

func main() {

    name := os.Getenv("DBNamedb")
    host := os.Getenv("HOSTdb")
    password := os.Getenv("PASSWORDdb")
    portdb := os.Getenv("PORTdb")
    
    db :=os.Getenv("USERdb")

    log.Println("DBnamedb:", name)
    log.Println("USERdb:", db)
    log.Println("HOSTdb:", host)
    log.Println("PASSWORDdb:", password)
    log.Println("PORTdb:", portdb)
    // Define config file path
    configFile := flag.String("config", "/etc/secrets/.env", "Path to configuration file")
    
    flag.Parse()

    // Try to load .env file, but don't fail if it doesn't exist (for Render deployment)
    _ = godotenv.Load(*configFile)
    
    // Запуск приложения
    if err := app.Run(*configFile); err != nil {
        log.Fatalf("Failed to run application: %v", err)
    }
}