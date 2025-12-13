package routes

import (
	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, container *app.Container) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"lambda":  "runnung",
			"status":  "ok",
			"message": "HyperShop API is running",
		})
	})

	// Health Check Route (Keep-Warm ke liye)
	router.GET("/health", func(c *gin.Context) {
		// 1. DB check (Zaroori hai TiDB ko jagane ke liye)
		sqlDB, err := config.DB.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "db_connection_error"})
			return
		}

		// Lightweight Query - "SELECT 1" sabse fast hoti hai
		if err := sqlDB.Ping(); err != nil {
			// Agar connection toot gaya hai toh reconnect try karega (agar handler me logic hai)
			// Ya bas error return karega taaki agli baar fresh connect ho
			c.JSON(500, gin.H{"status": "db_ping_failed"})
			return
		}

		c.JSON(200, gin.H{
			"status":  "alive",
			"message": "Lambda and TiDB are warm!",
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
