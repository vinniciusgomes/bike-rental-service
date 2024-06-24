package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

func RentalHandler(router *gin.Engine, rentalService *services.RentalService) {
	v1 := router.Group("/v1")
	{
		rentalRouter := v1.Group("/rentals")
		rentalRouter.Use(middlewares.AuthMiddleware())
		{
			rentalRouter.POST("/rent/:bikeId", rentalService.CreateRental)
			rentalRouter.POST("/return/:rentalId", rentalService.ReturnBike)
			rentalRouter.GET("/:userId", rentalService.GetRentalByUserID)
		}
	}

	admin := router.Group("/v1/admin")
	{
		adminRouter := admin.Group("/rentals")
		adminRouter.Use(middlewares.AuthMiddleware())
		adminRouter.Use(middlewares.AdminOnly())
		{
			adminRouter.GET("/", rentalService.GetAllRentals)
		}
	}
}
