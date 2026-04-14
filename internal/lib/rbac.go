package lib

import (
	"koda-b6-backend1/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		roleVal, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "unauthorized",
			})
			c.Abort()
			return
		}

		role := roleVal.(string)

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, models.Response{
			Success: false,
			Message: "forbidden",
		})
		c.Abort()
	}
}
