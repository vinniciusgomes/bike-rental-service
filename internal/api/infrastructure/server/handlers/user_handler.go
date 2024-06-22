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
func UserHandler(router *gin.Engine, userService *services.UserService) {
	v1 := router.Group("/v1")
	{
		userRouter := v1.Group("/users")
		userRouter.Use(middlewares.AuthMiddleware())
		{
			userRouter.GET("/:id", userService.GetUserByID)
			userRouter.PUT("/:id", userService.UpdateUser)
			userRouter.PUT("/:id/password", userService.UpdatePassword)
			userRouter.DELETE("/:id/delete", userService.DeleteUser)
		}
	}

	adminGroup := router.Group("/v1/admin")
	adminGroup.Use(middlewares.AuthMiddleware(), middlewares.AdminOnly())
	{
		adminUserRouter := adminGroup.Group("/users")
		{
			adminUserRouter.GET("", userService.GetAllUsers)
		}
	}
}
