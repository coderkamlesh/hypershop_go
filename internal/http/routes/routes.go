package routes

import (
	"net/http"

	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/coderkamlesh/hypershop_go/internal/http/dto"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, container *app.Container) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, dto.Success("running", gin.H{
			"lambda":  "running",
			"status":  "ok",
			"message": "HyperShop API is running",
		}))
	})

	// Health Check Route (Keep-Warm ke liye)
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := config.DB.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "db_connection_error"})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "db_ping_failed"})
			return
		}

		c.JSON(200, gin.H{
			"status":  "alive",
			"message": "App and DB are warm!",
		})
	})

	// API v1 group
	v1 := router.Group("/api/v1")

	// Simple Image Upload Route
	v1.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, dto.Failure("please provide image."))
			return
		}
		defer file.Close()

		// Upload to gallery folder
		imageURL, err := utils.UploadFile(file, header.Filename, "hypershop/gallery")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"url":     imageURL,
		})
	})

	// Setup domain routes
	SetupAuthRoutes(v1, container)
	SetupUserRoutes(v1, container)
	SetupProductRoutes(v1, container)
	SetupOrderRoutes(v1, container)
}
