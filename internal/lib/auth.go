package lib

import (
	"koda-b6-backend1/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" || token != "1234" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Message: "Unauthorized",
			})
			return
		} else if 
		ctx.Next()
	}
}
