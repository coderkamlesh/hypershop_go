package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/coderkamlesh/hypershop_go/internal/http/routes"
	ginFramework "github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

// init runs once when Lambda cold starts
func init() {
	log.Println("🚀 Lambda Cold Start - Initializing...")

	// 1. Load environment variables
	config.LoadEnv()

	// 2. Set Gin to release mode for production
	ginFramework.SetMode(ginFramework.ReleaseMode)

	// 3. Connect to MongoDB
	config.ConnectDB()

	// 4. Create application container
	container := app.NewContainer()

	// 5. Setup Gin router
	router := ginFramework.Default()

	// Add CORS middleware for API Gateway
	router.Use(corsMiddleware())

	// Setup routes
	routes.SetupRoutes(router, container)

	// 6. Wrap Gin router with Lambda adapter
	ginLambda = ginadapter.New(router)

	log.Println("✅ Lambda initialization complete")
}

// Handler is the Lambda function handler
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// ⚠️ CHECK: Har request pe ensure karo DB zinda hai
	// TiDB connection kabhi kabhi drop ho jata hai idle rehne par
	sqlDB, err := config.DB.DB() // GORM se underlying SQL DB nikalo
	if err == nil {
		if err := sqlDB.Ping(); err != nil {
			log.Println("⚠️ DB Connection lost, reconnecting...")
			config.ConnectDB() // Reconnect agar ping fail ho
		}
	} else {
		// Agar sqlDB object hi nahi mila
		config.ConnectDB()
	}

	// Proxy request to Gin router
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

// corsMiddleware adds CORS headers
func corsMiddleware() ginFramework.HandlerFunc {
	return func(c *ginFramework.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
