package middleware

import (
	"net/http"

	"github.com/coderkamlesh/hypershop_go/internal/http/dto"
	"github.com/gin-gonic/gin"
)

func ConsumerLoginGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientType := c.GetHeader("X-Client-Type")

		if clientType != "consumer-app" {
			c.JSON(http.StatusUnauthorized, dto.Failure("Invalid client type"))
			c.Abort()
			return
		}
		c.Next()
	}
}
