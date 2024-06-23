package services

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/helpers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
)

type BikeService struct {
	repo repositories.BikeRepository
}

func NewBikeService(repo repositories.BikeRepository) *BikeService {
	return &BikeService{repo: repo}
}

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

	if err := helpers.ValidateModel(bike); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if err := s.repo.CreateBike(bike); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to create a new bike"})
		return
	}

	c.Status(http.StatusCreated)
}

func (s *BikeService) GetAllBikes(c *gin.Context) {
	bikes, err := s.repo.GetAllBikes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurrend when trying to get all bikes"})
		return
	}

	c.JSON(http.StatusOK, bikes)
}

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
