package main

import (
	"fmt"

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
	router := gin.Default()
	routes.SetupRoutes(router, container)

	// 6. Start server
	// port := fmt.Sprintf(":%s", config.AppConfig.Port)
	// fmt.Printf("🚀 Server running on http://localhost%s\n", port)
	// router.Run(port)

	address := fmt.Sprintf("0.0.0.0:%s", config.AppConfig.Port)
	fmt.Printf("🚀 Server running on http://%s\n", address)
	router.Run(address)

}
