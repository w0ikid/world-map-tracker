package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/w0ikid/world-map-tracker/internal/app/config"
	"github.com/w0ikid/world-map-tracker/internal/api/handlers"
	"github.com/w0ikid/world-map-tracker/internal/api/middlewares"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config, UserUseCase *usecase.UserUseCase) {
	userHandler := handlers.NewUserHandler(UserUseCase)

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
		}

	}
}