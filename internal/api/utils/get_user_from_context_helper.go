package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
)

// GetUserFromContext retrieves the user from the gin.Context and returns a pointer to the models.User and an error.
//
// Parameters:
// - c: a pointer to a gin.Context object representing the HTTP request and response.
//
// Returns:
// - *models.User
func GetUserFromContext(c *gin.Context) (*models.User, error) {
	user, exists := c.Get("user")
	if !exists {
		return nil, errors.New("user not found")
	}

	return user.(*models.User), nil
}
