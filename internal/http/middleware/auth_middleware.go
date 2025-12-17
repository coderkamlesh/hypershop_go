package middleware

import (
	"net/http"
	"strings"

	"github.com/coderkamlesh/hypershop_go/internal/constants"
	"github.com/coderkamlesh/hypershop_go/internal/http/dto"
	"github.com/coderkamlesh/hypershop_go/internal/service"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(
	authService service.AuthService,
	allowedRoles ...constants.Role,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		// ================= 1. Authorization Header =================
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.InvalidTokenWithMessage("authorization header missing"))
			c.Abort()
			return
		}

		// Expecting: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, dto.InvalidTokenWithMessage("invalid authorization format"))
			c.Abort()
			return
		}

		token := parts[1]

		// ================= 2. Validate Token + Session =================
		if err := authService.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, dto.InvalidTokenWithMessage(err.Error()))
			c.Abort()
			return
		}

		// ================= 3. Extract Claims =================
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.InvalidToken())
			c.Abort()
			return
		}

		role := constants.Role(claims.Role)

		// ================= 4. Validate Role =================
		if !role.IsValid() {
			c.JSON(http.StatusForbidden, dto.Failure("invalid role"))
			c.Abort()
			return
		}

		// ================= 5. Role Authorization =================
		if len(allowedRoles) > 0 {
			allowed := false
			for _, r := range allowedRoles {
				if r == role {
					allowed = true
					break
				}
			}

			if !allowed {
				c.JSON(http.StatusForbidden, dto.InvalidToken())
				c.Abort()
				return
			}
		}

		// ================= 6. Set Context =================
		c.Set("userId", claims.UserID)
		c.Set("role", role)
		c.Set("token", token)

		// ================= 7. Continue =================
		c.Next()
	}
}
