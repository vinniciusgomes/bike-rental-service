package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/infrastructure/server/middlewares"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

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
			adminRouter.POST("/new", bikeService.CreateBike)
		}
	}
}
