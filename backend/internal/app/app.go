package app

import (
	"fmt"
	"log"

	"github.com/w0ikid/world-map-tracker/internal/app/start"
	"github.com/w0ikid/world-map-tracker/internal/domain/repositories"
	"github.com/w0ikid/world-map-tracker/internal/domain/services"

	"github.com/w0ikid/world-map-tracker/internal/app/config"
	"github.com/w0ikid/world-map-tracker/internal/app/connections"
	"github.com/w0ikid/world-map-tracker/internal/app/migrations"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
)

func Run(configFile string) error {
	// conf load
	cfg, err := config.NewConfig(configFile)
	if err != nil {
		return err
	}
	log.Println("Config loaded successfully")
	// db connection


	conn, err := connections.NewConnections(cfg)
	if err != nil {
		return err
	}
	defer conn.Close()
	
	log.Println("DB connection established successfully")

	// auto migration
	if err := migrations.AutoMigrate(conn.DB); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	// load countries
	if err := migrations.LoadCountriesFromFile(conn.DB); err != nil {
		return fmt.Errorf("failed to load countries: %w", err)
	}
	
	userRepo := repositories.NewUserRepository(conn.DB)
	countryStatus := repositories.NewCountryStatusesRepository(conn.DB)

	userService := services.NewUserService(userRepo)
	countryStatusesService := services.NewCountryStatusesService(countryStatus)

	userUseCase := usecase.NewUserUseCase(userService)
	countryStatusesUseCase := usecase.NewCountryStatusesUseCase(countryStatusesService)
	start.HTTP(cfg, userUseCase, countryStatusesUseCase)
	
	return nil
}
