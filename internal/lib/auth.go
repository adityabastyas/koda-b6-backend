package lib

import (
	"koda-b6-backend1/internal/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		claims, err := ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}

		// fmt.Println(claims["user_id"])

		ctx.Set("user_id", int(claims["user_id"].(float64)))
		ctx.Next()

	}
}
