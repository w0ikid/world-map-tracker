package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/w0ikid/world-map-tracker/internal/app/config"
	"github.com/w0ikid/world-map-tracker/internal/api/handlers"
	"github.com/w0ikid/world-map-tracker/internal/api/middlewares"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, userUseCase *usecase.UserUseCase, countryStatusUseCase *usecase.CountryStatusesUseCase) {
	userHandler := handlers.NewUserHandler(userUseCase)
	countryStatus := handlers.NewCountryStatusesHandler(countryStatusUseCase)

	authMiddleware := middlewares.AuthMiddleware()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register",  userHandler.CreateUser)
			auth.POST("/login", userHandler.LoginUser)
			auth.POST("/logout", userHandler.LogoutUser)
		}
		users := api.Group("/users")
		users.Use(authMiddleware)
		{
			users.GET("/profile", userHandler.Profile)
			users.GET("/:username", userHandler.UserByUsername)
		}
		countries := api.Group("/countries")
		countries.Use(authMiddleware)
		{
			countries.POST("/", countryStatus.CreateCountryStatus)
			countries.GET("/", countryStatus.GetCountryStatuses)
			countries.PUT("/", countryStatus.UpdateCountryStatus)
			countries.DELETE("/", countryStatus.DeleteCountryStatus)
			countries.GET("/visited-percentage", countryStatus.GetVisitedPercentage)
			countries.GET("/visited-count", countryStatus.GetVisitedCount)
		}
	}
}