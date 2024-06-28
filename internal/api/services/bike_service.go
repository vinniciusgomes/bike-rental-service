package services

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/utils"
)

type BikeService struct {
	repo repositories.BikeRepository
}

// NewBikeService creates a new instance of the BikeService struct.
//
// It takes a repositories.BikeRepository as a parameter and returns a pointer
// to a BikeService.
func NewBikeService(repo repositories.BikeRepository) *BikeService {
	return &BikeService{repo: repo}
}

// CreateBike creates a new bike based on the JSON input in the request body.
//
// It uses the gin.Context parameter to access the HTTP request and response.
// Returns nothing.
func (s *BikeService) CreateBike(c *gin.Context) {
	bike := new(models.Bike)

	id, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to create user"})
		return
	}

	if err := c.BindJSON(&bike); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid body"})
		return
	}

	bike.ID = id

	if err := utils.ValidateModel(bike); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := s.repo.CreateBike(bike); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to create a new bike"})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllBikes retrieves all the bikes from the repository and returns them in the response.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
//
// Return:
// - None.
func (s *BikeService) GetAllBikes(c *gin.Context) {
	limit, offset := utils.GetPaginationParams(c)

	bikes, err := s.repo.GetAllBikes(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to get all bikes"})
		return
	}

	c.JSON(http.StatusOK, bikes)
}

// GetBikeByID retrieves a bike by its ID from the BikeService repository and returns it in the response.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
//
// Return:
// - None.
func (s *BikeService) GetBikeByID(c *gin.Context) {
	id := c.Param("id")

	bike, err := s.repo.GetBikeByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "bike not found") {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to get bike"})
		return
	}

	c.JSON(http.StatusOK, bike)
}

// DeleteBike deletes a bike identified by ID from the BikeService repository.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
// Return:
// - None.
func (s *BikeService) DeleteBike(c *gin.Context) {
	id := c.Param("id")

	err := s.repo.DeleteBike(id)
	if err != nil {
		if strings.Contains(err.Error(), "bike not found") {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to delete bike"})
		return
	}

	c.Status(http.StatusOK)
}

// UpdateBike updates a bike entity based on the input JSON data in the request body.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
// Return:
// - None.
func (s *BikeService) UpdateBike(c *gin.Context) {
	id := c.Param("id")

	var bike models.Bike
	if err := c.BindJSON(&bike); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bikeID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to update bike"})
		return
	}

	bike.ID = bikeID

	if err := s.repo.UpdateBike(&bike); err != nil {
		if strings.Contains(err.Error(), "bike not found") {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to update bike"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bike updated successfully"})
}
