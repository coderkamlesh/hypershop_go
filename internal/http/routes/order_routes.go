package routes

import (
	"net/http"

	"github.com/coderkamlesh/hypershop_go/internal/app"
	"github.com/coderkamlesh/hypershop_go/internal/constants"
	"github.com/coderkamlesh/hypershop_go/internal/http/dto"
	"github.com/coderkamlesh/hypershop_go/internal/http/middleware"
	"github.com/gin-gonic/gin"
)

// SetupOrderRoutes configures order routes
func SetupOrderRoutes(v1 *gin.RouterGroup, container *app.Container) {

	orders := v1.Group("/orders")

	// ================= PUBLIC (NO AUTH) =================
	orders.GET("/public-test", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.SuccessWithoutData("public order api working"))
	})

	// ================= CONSUMER ONLY =================
	orders.GET(
		"/consumer-test",
		middleware.AuthMiddleware(
			container.AuthService,
			constants.RoleConsumer,
		),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, dto.SuccessWithoutData("consumer order api working"))
		},
	)

	// ================= ADMIN + SELLER =================
	orders.GET(
		"/admin-seller-test",
		middleware.AuthMiddleware(
			container.AuthService,
			constants.RoleAdmin,
			constants.RoleSeller,
		),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, dto.SuccessWithoutData("admin/seller order api working"))
		},
	)
}
