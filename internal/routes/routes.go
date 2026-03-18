package routes

import (
	"koda-b6-backend1/internal/handlers"
	"koda-b6-backend1/internal/lib"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler, kategoryHandler *handlers.KategoryHandler, promoHandler *handlers.PromoHandler, discountHandler *handlers.DiscountHandler, cartHandler *handlers.CartHandler) {

	r.Use(lib.CorsMiddleware())

	authGroup := r.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
	authGroup.POST("/forgot-password", authHandler.RequestForgotPassword)
	authGroup.PATCH("/forgot-password", authHandler.ResetPassword)

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

	//kategory
	kategoryGroup := r.Group("/kategorys")
	kategoryGroup.GET("", kategoryHandler.GetAll)
	kategoryGroup.GET("/:id", kategoryHandler.GetByID)
	kategoryGroup.Use(lib.AuthMiddleware())
	kategoryGroup.POST("", kategoryHandler.Create)
	kategoryGroup.PATCH("/:id", kategoryHandler.Update)
	kategoryGroup.DELETE("/:id", kategoryHandler.Delete)

	// promo routes
	promoGroup := r.Group("/promos")
	promoGroup.GET("", promoHandler.GetAll)
	promoGroup.GET("/:id", promoHandler.GetByID)
	promoGroup.Use(lib.AuthMiddleware())
	promoGroup.POST("", promoHandler.Create)
	promoGroup.PATCH("/:id", promoHandler.Update)
	promoGroup.DELETE("/:id", promoHandler.Delete)

	//discount
	discountGroup := r.Group("/discounts")
	discountGroup.GET("", discountHandler.GetAll)
	discountGroup.GET("/:id", discountHandler.GetByID)
	discountGroup.Use(lib.AuthMiddleware())
	discountGroup.POST("", discountHandler.Create)
	discountGroup.PATCH("/:id", discountHandler.Update)
	discountGroup.DELETE("/:id", discountHandler.Delete)

	//cart
	cartGroup := r.Group("/carts")
	cartGroup.Use(lib.AuthMiddleware())
	cartGroup.GET("", cartHandler.GetAll)
	cartGroup.GET("/:user_id", cartHandler.GetByUserID)

}
