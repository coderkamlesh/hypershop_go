package routes

import (
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine, container *app.Container) {
	// Health check endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"lambda":  "runnung",
			"status":  "ok",
			"message": "HyperShop API is running",
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")

	// Setup domain routes
	SetupAuthRoutes(v1, container)
	SetupUserRoutes(v1, container)
	SetupProductRoutes(v1, container)
	SetupOrderRoutes(v1, container)
}
