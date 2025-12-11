package routes

import (
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/gin-gonic/gin"
)

// SetupOrderRoutes configures order routes
func SetupOrderRoutes(v1 *gin.RouterGroup, container *app.Container) {
	// orders := v1.Group("/orders")
	// {
	// TODO: Uncomment when OrderHandler is ready
	// orders.GET("", container.OrderHandler.GetAllOrders)
	// orders.GET("/:id", container.OrderHandler.GetOrder)
	// orders.POST("", container.OrderHandler.CreateOrder)
	// orders.PUT("/:id/status", container.OrderHandler.UpdateOrderStatus)
	// }
}
