package routes

import (
	"koda-b6-backend1/internal/handlers"
	"koda-b6-backend1/internal/lib"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler) {
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	userGroup := r.Group("/users")
	userGroup.Use(lib.AuthMiddleware())

	userGroup.GET("", userHandler.GetAll)
}
