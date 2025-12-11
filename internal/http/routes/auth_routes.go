package routes

import (
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configures all authentication routes
func SetupAuthRoutes(v1 *gin.RouterGroup, container *app.Container) {
	auth := v1.Group("/auth")

	// ========== CONSUMER REGISTRATION ==========
	consumer := auth.Group("/consumer")
	{
		consumer.POST("/register/requestOtp", container.AuthHandler.RequestConsumerRegistrationOTP)
		consumer.POST("/register/verifyOtp", container.AuthHandler.VerifyConsumerRegistrationOTP)
		consumer.POST("/login/requestOtp", container.AuthHandler.RequestConsumerLoginOTP)
		consumer.POST("/login/verifyOtp", container.AuthHandler.VerifyConsumerLoginOTP)
	}

	// ========== ADMIN LOGIN ==========
	admin := auth.Group("/admin")
	{
		admin.POST("/login/requestOtp", container.AuthHandler.RequestAdminLoginOTP)
		admin.POST("/login/verifyOtp", container.AuthHandler.VerifyAdminLoginOTP)
	}

	// ========== RIDER LOGIN ==========
	rider := auth.Group("/rider")
	{
		rider.POST("/login/requestOtp", container.AuthHandler.RequestRiderLoginOTP)
		rider.POST("/login/verifyOtp", container.AuthHandler.VerifyRiderLoginOTP)
	}

	// ========== MANAGER LOGIN ==========
	manager := auth.Group("/manager")
	{
		manager.POST("/login/requestOtp", container.AuthHandler.RequestManagerLoginOTP)
		manager.POST("/login/verifyOtp", container.AuthHandler.VerifyManagerLoginOTP)
	}

	// ========== TOKEN VALIDATION ==========
	auth.GET("/validate-token", container.AuthHandler.ValidateToken)
}
