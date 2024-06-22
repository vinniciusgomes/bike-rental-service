package services

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/constants"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/helpers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositories.UserRepository
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Image     string    `json:"image"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
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

// GetAllUsers retrieves users based on specified filters from the request query parameters.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
//
// Returns:
// This function does not return anything. It sends a JSON response with the user information if successful or an error response if an error occurs.
func (s *UserService) GetAllUsers(c *gin.Context) {
	filters := make(map[string]interface{})

	if id := strings.TrimSpace(c.Query("id")); id != "" {
		filters["id"] = id
	} else if email := strings.TrimSpace(c.Query("email")); email != "" {
		filters["email"] = email
	} else if name := strings.TrimSpace(c.Query("name")); name != "" {
		filters["name"] = name
	} else if status := strings.TrimSpace(c.Query("status")); status != "" {
		if status != models.UserStatusActive && status != models.UserStatusInactive {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid status"})
			return
		}

		filters["status"] = status
	} else if role := strings.TrimSpace(c.Query("role")); role != "" {
		if role != models.UserRoleAdmin && role != models.UserRoleDefault {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid role"})
			return
		}

		filters["role"] = role
	}

	users, err := s.repo.GetAllUsers(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get users"})
		return
	}

	response := []UserResponseDTO{}
	for _, user := range *users {
		response = append(response, UserResponseDTO{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			Phone:     user.Phone,
			Image:     user.Image,
			Status:    user.Status,
			Role:      user.Role,
			Verified:  user.Verified,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
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
		Phone:     user.Phone,
		Image:     user.Image,
		Status:    user.Status,
		Role:      user.Role,
		Verified:  user.Verified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates a user's information.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
//
// Return type: None.
func (s *UserService) UpdateUser(c *gin.Context) {
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

	var body struct {
		Email string `json:"email" validate:"required,email"`
		Name  string `json:"name" validate:"required"`
		Phone string `json:"phone" validate:"required"`
		Image string `json:"image"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	if err := helpers.ValidateModel(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = s.repo.UpdateUser(&models.User{
		ID:    uuid.MustParse(id),
		Email: body.Email,
		Name:  body.Name,
		Phone: body.Phone,
		Image: body.Image,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

// UpdatePassword updates the user's password.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
// Return type: None.
func (s *UserService) UpdatePassword(c *gin.Context) {
	id := c.Params.ByName("id")
	loggedUser, err := helpers.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
		return
	}

	if loggedUser.ID.String() != id {
		c.JSON(http.StatusForbidden, gin.H{"message": "access to this resource is forbidden"})
		return
	}

	var body struct {
		CurrentPassword string `json:"current_password" validate:"required"`
		NewPassword     string `json:"password" validate:"required"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	if err := helpers.ValidateModel(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	if body.CurrentPassword == body.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "new password cannot be the same as current password"})
		return
	}

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.CurrentPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid current password"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to update password"})
		return
	}

	err = s.repo.UpdateUser(&models.User{
		ID:       uuid.MustParse(id),
		Password: string(hashedPassword),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated successfully"})
}

// DeleteUser deletes a user from the system.
//
// Parameters:
// - c: The gin.Context object representing the HTTP request and response.
//
// Return type: None.
func (s *UserService) DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	loggedUser, err := helpers.GetUserFromContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
		return
	}

	if loggedUser.ID.String() != id {
		c.JSON(http.StatusForbidden, gin.H{"message": "access to this resource is forbidden"})
		return
	}

	err = s.repo.DeleteUser(loggedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to delete user"})
		return
	}

	c.SetCookie(constants.AuthCookieName, "", -1, "", "", false, true)
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
