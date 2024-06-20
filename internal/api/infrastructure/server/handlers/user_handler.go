package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

// UserHandler registers the user handler with the provided gin.Engine and services.UserService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - UserService: a pointer to a services.UserService object providing the user-related operations.
func UserHandler(router *gin.Engine, UserService *services.UserService) {
	v1 := router.Group("/v1")
	{
		userRouter := v1.Group("/users")
		{
			userRouter.GET("/:id", middlewares.AuthMiddleware(), UserService.GetUserByID)
		}
	}
}
