package server

import (
	"log"
	"math"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/config"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/handlers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

// isPrime verifica se um número é primo.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// generatePrimes gera números primos dentro de um intervalo e os adiciona a um slice.
func generatePrimes(start, end int, wg *sync.WaitGroup, primes *[]int) {
	defer wg.Done() // Marca o trabalho como concluído quando a função termina.
	for i := start; i <= end; i++ {
		if isPrime(i) {
			*primes = append(*primes, i) // Adiciona o número primo ao slice.
		}
	}
}

// cpuIntensiveTask é o handler que executa a tarefa intensiva de CPU.
func cpuIntensiveTask(c *gin.Context) {
	numWorkers := 4                        // Número de goroutines (trabalhadores) a serem usadas.
	rangePerWorker := 1000000 / numWorkers // Intervalo de números para cada goroutine.
	var wg sync.WaitGroup
	primes := make([]int, 0) // Slice para armazenar os números primos.

	// Inicia as goroutines para gerar números primos.
	for i := 0; i < numWorkers; i++ {
		start := i*rangePerWorker + 1
		end := (i + 1) * rangePerWorker
		wg.Add(1) // Incrementa o contador do WaitGroup.
		go generatePrimes(start, end, &wg, &primes)
	}

	wg.Wait() // Espera todas as goroutines terminarem.

	// Retorna a resposta em JSON.
	c.JSON(http.StatusOK, gin.H{
		"message": "Tarefa intensiva de CPU concluída",
		"primes":  primes,
	})
}

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

	router.GET("/cpu-intensive", cpuIntensiveTask)

	// Start server
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	return router.Run(":" + httpPort)
}
