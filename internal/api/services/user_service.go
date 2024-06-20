package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/helpers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	Image     *string   `json:"image"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUserService creates a new instance of the UserService struct.
//
// It takes a repositories.UserRepository as a parameter and returns a pointer
// to a UserService.
func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetUserByID retrieves a user by their ID.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
//
// Returns:
// This function does not return anything. It sends a JSON response with the user's information if the user is found,
// or a JSON error response if the user is not found or an error occurs.
func (s *UserService) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	loggedUser, err := helpers.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
		return
	}

	if loggedUser.Role != models.UserRoleAdmin && loggedUser.ID.String() != id {
		c.JSON(http.StatusForbidden, gin.H{"message": "access to this resource is forbidden"})
		return
	}

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	response := UserResponseDTO{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Status:    user.Status,
		Role:      user.Role,
		Image:     user.Image,
		Verified:  user.Verified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
