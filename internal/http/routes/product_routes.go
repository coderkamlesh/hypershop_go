package routes

import (
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/gin-gonic/gin"
)

// SetupProductRoutes configures product routes
func SetupProductRoutes(v1 *gin.RouterGroup, container *app.Container) {
	// products := v1.Group("/products")
	// {
	// TODO: Uncomment when ProductHandler is ready
	// products.GET("", container.ProductHandler.GetAllProducts)
	// products.GET("/:id", container.ProductHandler.GetProduct)
	// products.POST("", container.ProductHandler.CreateProduct)
	// products.PUT("/:id", container.ProductHandler.UpdateProduct)
	// products.DELETE("/:id", container.ProductHandler.DeleteProduct)
	// }
}
