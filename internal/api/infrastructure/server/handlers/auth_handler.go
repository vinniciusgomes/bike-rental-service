package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinniciusgomes/ebike-rental-service/internal/api/services"
)

// AuthHandler registers the auth handler with the provided gin.Engine and services.AuthService.
//
// Parameters:
// - router: a pointer to a gin.Engine object representing the HTTP router.
// - authService: a pointer to a services.AuthService object providing the auth-related operations.
//
// Returns: None.
func AuthHandler(router *gin.Engine, authService *services.AuthService) {
	v1 := router.Group("/v1")
	{
		authRouter := v1.Group("/auth")
		{
			authRouter.POST("/register", authService.CreateUser)
			authRouter.POST("/login", authService.Login)
			authRouter.POST("/logout", authService.Logout)
			authRouter.POST("/verify-account", authService.VerifyUser)           // TODO
			authRouter.POST("/forgot-password", authService.ForgotPassword)      // TODO
			authRouter.POST("/reset-password/:token", authService.ResetPassword) // TODO
			authRouter.POST("/refresh-token", authService.RefreshToken)          // TODO
		}
	}
}
