package services

import (
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/helpers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
)

type RentalService struct {
	repo repositories.RentalRepository
}

// NewRentalService creates a new instance of the RentalService struct.
//
// It takes a repositories.RentalRepository as a parameter and returns a pointer
// to a RentalService.
//
// Parameters:
// - repo: a repositories.RentalRepository object representing the rental repository.
//
// Returns:
// - *RentalService: a pointer to a RentalService object.
func NewRentalService(repo repositories.RentalRepository) *RentalService {
	return &RentalService{repo: repo}
}

// CreateRental creates a new rental for a bike.
//
// It takes a gin.Context object as a parameter and returns nothing.
//
// Parameters:
// - c: a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None.
func (s *RentalService) CreateRental(c *gin.Context) {
	bikeID := c.Param("bikeId")

	loggedUser, err := helpers.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
		return
	}

	bike, err := s.repo.GetBikeByID(bikeID)
	if err != nil {
		if strings.Contains(err.Error(), "bike not found") {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to rent a bike"})
		return
	}

	if bike.Status != models.BIKE_STATUS_AVAILABLE {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bike is not available to rent"})
		return
	}

	rental := &models.Rental{
		ID:        uuid.Must(uuid.NewRandom()),
		UserID:    loggedUser.ID,
		BikeID:    bike.ID,
		StartTime: time.Now(),
		Status:    models.RENTAL_STATUS_ACTIVE,
	}

	if err := helpers.ValidateModel(rental); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := s.repo.CreateRental(rental); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to create a new rental"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rental_id": rental.ID})
}

// ReturnBike handles the process of returning a rented bike.
//
// Parameters:
// - c: a gin.Context object representing the HTTP request and response.
// Return type(s): None.
func (s *RentalService) ReturnBike(c *gin.Context) {
	id := c.Param("rentalId")

	loggedUser, err := helpers.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
		return
	}

	rental, err := s.repo.GetRentalByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "rental not found") {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to retrieve rental"})
		return
	}

	if loggedUser.Role != models.UserRoleAdmin && rental.UserID != loggedUser.ID {
		c.JSON(http.StatusForbidden, gin.H{"message": "you are not allowed to complete this rental"})
		return
	}

	if rental.Status == models.RENTAL_STATUS_COMPLETED {
		c.JSON(http.StatusBadRequest, gin.H{"message": "rental already completed"})
		return
	}

	bike, err := s.repo.GetBikeByID(rental.BikeID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to retrieve bike"})
		return
	}

	now := time.Now()
	rental.EndTime = now
	rental.Status = models.RENTAL_STATUS_COMPLETED

	duration := rental.EndTime.Sub(rental.StartTime).Hours()
	totalCost := math.Round(duration*bike.PricePerHour*100) / 100
	rental.TotalCost = totalCost

	if err := s.repo.UpdateRental(rental); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to update rental"})
		return
	}

	if err := s.repo.UpdateBikeStatus(bike.ID.String(), models.BIKE_STATUS_AVAILABLE); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to update bike status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_time":  duration,
		"total_price": totalCost,
	})
}

// GetAllRentals retrieves all rentals.
//
// Parameters:
// - c: a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None.
func (s *RentalService) GetAllRentals(c *gin.Context) {
	rentals, err := s.repo.GetAllRentals()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get rentals"})
		return
	}

	c.JSON(http.StatusOK, rentals)
}

// GetRentalByUserID retrieves rentals for a specific user.
//
// It takes a gin.Context object as a parameter and returns nothing.
//
// Parameters:
// - c: a gin.Context object representing the HTTP request and response.
//
// Returns:
// - None.
func (s *RentalService) GetRentalByUserID(c *gin.Context) {
	id := c.Param("userId")

	rentals, err := s.repo.GetRentalByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get rentals"})
		return
	}

	c.JSON(http.StatusOK, rentals)
}
