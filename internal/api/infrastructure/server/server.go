package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
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
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

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