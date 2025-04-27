package start

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/w0ikid/world-map-tracker/internal/domain/usecase"
	"github.com/w0ikid/world-map-tracker/internal/app/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
	"github.com/w0ikid/world-map-tracker/internal/api/routes"
)

func HTTP(cfg *config.Config, userUseCase *usecase.UserUseCase, countryStatusesUseCase *usecase.CountryStatusesUseCase) {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // Установите true, если используете HTTPS
	})

	router.Use(sessions.Sessions("session", store))


	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "https://world-map-tracker-nine.vercel.app/"}, // адрес фронта
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	routes.SetupRoutes(router, cfg, userUseCase, countryStatusesUseCase)


	srv := &http.Server{
		Addr:    ":" + cfg.HTTPServer.Port,
		Handler: router,
	}
	log.Printf("Starting server on port %s", cfg.HTTPServer.Port)
	// Запускаем сервер в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Ждем сигнала для грациозного завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Даем 5 секунд на завершение текущих запросов
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}