package routes

import (
	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures user management routes
func SetupUserRoutes(v1 *gin.RouterGroup, container *app.Container) {
	// users := v1.Group("/users")
	{
		// users.GET("", container.UserHandler.GetAllUsers)
		// users.GET("/:id", container.UserHandler.GetUser)
		// users.PUT("/:id", container.UserHandler.UpdateUser)
		// users.DELETE("/:id", container.UserHandler.DeleteUser)
	}
}
