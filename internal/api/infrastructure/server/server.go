package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/config"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/handlers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

// StartServer starts the server with the provided configuration.
//
// It checks if the .env file exists and loads it if it does. Then, it creates
// a Gin router with default middleware and sets up the "/health" route.
// Finally, it starts the server on the specified port, or "8080" if no port
// is provided.
//
// Returns an error if there was a problem loading the .env file or starting
// the server.
func StartServer() error {
	// Check if .env file exists
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Database
	config.DatabaseInit()
	gorm := config.GetDatabaseInstance()
	dbGorm, err := gorm.DB()
	if err != nil {
		return err
	}

	err = dbGorm.Ping()
	if err != nil {
		panic(err)
	}
	defer dbGorm.Close()

	// Init router
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	// Services
	authService := services.NewAuthService(repositories.NewAuthRepository(config.GetDatabaseInstance()))
	userService := services.NewUserService(repositories.NewUserRepository(config.GetDatabaseInstance()))
	bikeService := services.NewBikeService(repositories.NewBikeRepository(config.GetDatabaseInstance()))
	rentalService := services.NewRentalService(repositories.NewRentalRepository(config.GetDatabaseInstance()))

	// Routes
	handlers.AuthHandler(router, authService)
	handlers.UserHandler(router, userService)
	handlers.BikeHandler(router, bikeService)
	handlers.RentalHandler(router, rentalService)

	// Others routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	// Start server
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return router.Run(":" + httpPort)
}
