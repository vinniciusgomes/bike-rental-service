package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

// BikeHandler handles HTTP requests related to bikes using the provided router and BikeService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - bikeService: a pointer to a services.BikeService object providing the bike-related operations.
func BikeHandler(router *gin.Engine, bikeService *services.BikeService) {
	v1 := router.Group("/v1")
	{
		bikesRouter := v1.Group("/bikes")
		bikesRouter.Use(middlewares.AuthMiddleware())
		{
			bikesRouter.GET("/", bikeService.GetAllBikes)
			bikesRouter.GET("/:id", bikeService.GetBikeByID)
		}
	}

	admin := router.Group("/v1/admin")
	{
		adminRouter := admin.Group("/bikes")
		adminRouter.Use(middlewares.AuthMiddleware())
		adminRouter.Use(middlewares.AdminOnly())
		{
			adminRouter.POST("/", bikeService.CreateBike)
			adminRouter.PUT("/:id", bikeService.UpdateBike)
			adminRouter.DELETE("/:id", bikeService.DeleteBike)
		}
	}
}
