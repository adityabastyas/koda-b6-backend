package routes

import (
	"koda-b6-backend1/internal/handlers"
	"koda-b6-backend1/internal/lib"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler) {
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	userGroup := r.Group("/users")
	userGroup.Use(lib.AuthMiddleware())
	userGroup.GET("", userHandler.GetAll)
	userGroup.POST("/upload", userHandler.UploadPhoto)

	//product
	productGroup := r.Group("/products")

	productGroup.GET("", productHandler.GetAll)
	productGroup.GET("/:id", productHandler.GetByID)

	productGroup.Use(lib.AuthMiddleware())
	productGroup.POST("", productHandler.Create)
	productGroup.PATCH("/:id", productHandler.Update)
	productGroup.DELETE("/:id", productHandler.Delete)
}
