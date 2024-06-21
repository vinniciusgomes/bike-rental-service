package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/helpers"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/models"
)

// AdminOnly is a middleware function for Gin framework that restricts access to
// routes to users with the admin role. It retrieves the logged-in user from the
// context, checks if the user has the admin role, and if not, returns an error
// response. If the user is an admin, it proceeds to the next handler in the
// chain.
//
// Parameters:
// - c: a pointer to the gin.Context object for handling HTTP request and response.
//
// Returns: void
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the user from the context
		loggedUser, err := helpers.GetUserFromContext(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occurred when trying to get user"})
			c.Abort()
			return
		}

		// Check if the user is an admin
		if loggedUser.Role != models.UserRoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"message": "access to this resource is forbidden"})
			c.Abort()
			return
		}

		// Proceed to the next handler
		c.Next()
	}
}
