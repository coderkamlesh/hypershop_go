// cmd/api/main.go
package main

import (
	"fmt"

	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
    // 1. Load environment variables
    config.LoadEnv()

    // 2. Set Gin mode
    gin.SetMode(config.AppConfig.GinMode)

    // 3. Connect MongoDB
    config.ConnectDB()

    // 4. Setup router
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status":  "ok",
            "message": "HyperShop API is running",
            "db":      config.AppConfig.DBName,
        })
    })

    // User routes
    v1 := router.Group("/api/v1")
    {
        userHandler := handlers.NewUserHandler()
        v1.POST("/users", userHandler.Create)
        v1.GET("/users", userHandler.GetAll)
        v1.GET("/users/:id", userHandler.GetByID)
        v1.PUT("/users/:id", userHandler.Update)
        v1.DELETE("/users/:id", userHandler.Delete)
    }

    // Server start
    port := fmt.Sprintf(":%s", config.AppConfig.Port)
    fmt.Printf("🚀 Server running on http://localhost%s\n", port)
    router.Run(port)
}
