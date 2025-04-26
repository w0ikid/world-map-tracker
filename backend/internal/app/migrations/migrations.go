package migrations

import (
	"fmt"
	"log"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"encoding/json"
	"github.com/w0ikid/world-map-tracker/internal/domain/models"
)

func AutoMigrate(conn *pgxpool.Pool) error {
	ctx := context.Background()
	log.Println("Starting auto migration...")
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(100) UNIQUE NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		);`,
		`CREATE TABLE IF NOT EXISTS countries (
			iso_code VARCHAR(2) PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS country_statuses (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			country_iso CHAR(2) NOT NULL REFERENCES countries(iso_code) ON DELETE RESTRICT,
			status VARCHAR(20) CHECK (status IN ('visited', 'wishlist', 'none')) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW(),
			UNIQUE (user_id, country_iso)
		);`,
	}

	for i, query := range queries {
		log.Printf("Executing query %d...\n", i+1)
		_, err := conn.Exec(ctx, query)
		if err != nil {
			fmt.Printf("Error executing query %d: %v\n", i+1, err)
			return fmt.Errorf("error executing query: %v", err)
		}
		log.Printf("Query %d executed successfully\n", i+1)
	}

	log.Println("Auto migration completed successfully")

	return nil
}

func LoadCountriesFromFile(conn *pgxpool.Pool) error {
	file, err := os.Open("internal/app/migrations/iso.json")
	if err != nil {
		return fmt.Errorf("failed to open countries file: %v", err)
	}
	defer file.Close()

	var countries []models.Country
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&countries); err != nil {
		return fmt.Errorf("failed to decode countries file: %v", err)
	}

	log.Printf("Loaded %d countries", len(countries))

	for _, country := range countries {
		_, err := conn.Exec(context.Background(), "INSERT INTO countries (iso_code, name) VALUES ($1, $2) ON CONFLICT (iso_code) DO NOTHING", country.ISO, country.Name)
		if err != nil {
			log.Printf("failed to insert country %s: %v", country.Name, err)
		}
	}

	return nil
}