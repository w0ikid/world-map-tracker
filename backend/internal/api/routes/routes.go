package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/w0ikid/world-map-tracker/internal/app/config"
	"github.com/w0ikid/world-map-tracker/internal/api/handlers"
	"github.com/w0ikid/world-map-tracker/internal/api/middlewares"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
	"github.com/w0ikid/world-map-tracker/internal/domain/services/llm"
	"os"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, userUseCase *usecase.UserUseCase, countryStatusUseCase *usecase.CountryStatusesUseCase) {
	userHandler := handlers.NewUserHandler(userUseCase)
	countryStatus := handlers.NewCountryStatusesHandler(countryStatusUseCase)

	authMiddleware := middlewares.AuthMiddleware()

	groqClient := llm.NewGroqClient(os.Getenv("GROQ_API_KEY"))
	groqHandler := handlers.NewLLMHandler(groqClient)

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
			users.GET("/interests", countryStatus.GetUsersWithSimilarList)
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
			countries.GET("/wish-list-count", countryStatus.GetWishListCount)
		}
		groq := api.Group("/ai")
		groq.Use(authMiddleware)
		{
			groq.GET("/ask", groqHandler.Ask)	
		}
		statistics := api.Group("/statistics")
		statistics.Use(authMiddleware)
		{
			statistics.GET("/top-visited", countryStatus.GetTopFiveVisitedCountries)
			statistics.GET("/top-wish-list", countryStatus.GetTopFiveWishlistCountries)
		}
	}
}