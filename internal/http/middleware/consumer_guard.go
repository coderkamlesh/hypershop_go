package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsumerLoginGuard() gin.HandlerFunc {
	return func(c *gin.Context) {

		clientType := c.GetHeader("X-Client-Type")

		if clientType != "consumer-app" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid client type",
			})
			c.Abort() // 🔴 chain yahin stop
			return
		}

		// ✅ sab theek hai → aage jaane do
		c.Next()
	}
}
