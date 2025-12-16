package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/coderkamlesh/hypershop_go/internal/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load environment variables
	config.LoadEnv()

	// 2. Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// 3. Connect Database
	config.ConnectDB()

	// 4. Create container (all handlers)
	container := app.NewContainer()

	// 5. Setup router
	router := gin.New()
	router.Use(gin.Recovery()) // ❗ recommended
	routes.SetupRoutes(router, container)

	// 6. Create HTTP server
	port := fmt.Sprintf(":%s", config.AppConfig.Port)
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// 7. Start server (non-blocking)
	go func() {
		fmt.Printf("🚀 Server running on http://localhost%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ listen error: %s\n", err)
		}
	}()

	// 8. Listen for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n🛑 Shutdown signal received")

	// 9. Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	// 10. Cleanup resources
	fmt.Println("🧹 Cleaning up resources...")

	// Close database connection
	if config.DB != nil {
		sqlDB, _ := config.DB.DB()
		if sqlDB != nil {
			sqlDB.Close()
			fmt.Println("✅ Database connection closed")
		}
	}
	fmt.Println("✅ Server exited gracefully")

}
