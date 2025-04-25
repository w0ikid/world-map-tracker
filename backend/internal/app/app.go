package app

import (
	"fmt"

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
	// db connection
	conn, err := connections.NewConnections(cfg)
	if err != nil {
		return err
	}
	defer conn.Close()

	// auto migration
	if err := migrations.AutoMigrate(conn.DB); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	// load countries
	if err := migrations.LoadCountriesFromFile(conn.DB); err != nil {
		return fmt.Errorf("failed to load countries: %w", err)
	}
	
	userRepo := repositories.NewUserRepository(conn.DB)


	userService := services.NewUserService(userRepo)


	userUseCase := usecase.NewUserUseCase(userService)

	start.HTTP(cfg, userUseCase)
	
	return nil
}
