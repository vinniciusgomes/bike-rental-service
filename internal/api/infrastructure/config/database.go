package config

import (
	"os"

	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database *gorm.DB

var database *gorm.DB
var err error

// GetDatabase initializes a connection to the PostgreSQL database specified by the
// "DATABASE_URL" environment variable and performs automatic migration of the
// "User" model.
//
// It does not take any parameters.
// It does not return anything.
// If there is an error opening the database connection, it panics.
func DatabaseInit() {
	database, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&models.ValidationToken{}, &models.User{}, &models.Bike{}, &models.Rental{})
	if err != nil {
		panic(err)
	}
}

// GetDatabaseInstance returns the database instance.
//
// It does not take any parameters.
// Returns *gorm.DB.
func GetDatabaseInstance() *gorm.DB {
	return database
}
