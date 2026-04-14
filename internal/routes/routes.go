package routes

import (
	"koda-b6-backend1/internal/handlers"
	"koda-b6-backend1/internal/lib"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, productHandler *handlers.ProductHandler, kategoryHandler *handlers.KategoryHandler, promoHandler *handlers.PromoHandler, discountHandler *handlers.DiscountHandler, cartHandler *handlers.CartHandler, transactionHandler *handlers.TransactionHandler, cartItemHandler *handlers.CartItemHandler, transactionProductHandler *handlers.TransactionProductHandler, productVariantHandler *handlers.ProductVariantHandler, productSizeHandler *handlers.ProductSizeHandler, productImagesHandler *handlers.ProductImagesHandler, reviewsHandler *handlers.ReviewsHandler) {

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
	userGroup.PATCH("/profile", userHandler.UpdateProfile)

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

	// transaction
	transactionGroup := r.Group("transactions")
	transactionGroup.Use(lib.AuthMiddleware())
	transactionGroup.GET("", transactionHandler.GetAll)
	transactionGroup.GET("/:id", transactionHandler.GetByID)
	transactionGroup.GET("/user/:user_id", transactionHandler.GetByUserID)
	transactionGroup.POST("", transactionHandler.Create)
	transactionGroup.DELETE("/:id", transactionHandler.Delete)

	/// cart-item routes
	cartItemGroup := r.Group("/cart-items")
	cartItemGroup.Use(lib.AuthMiddleware())
	cartItemGroup.GET("/:user_id", cartItemHandler.GetByUserID)
	cartItemGroup.POST("/:user_id", cartItemHandler.Create)
	cartItemGroup.PATCH("/:id", cartItemHandler.Update)
	cartItemGroup.DELETE("/:id", cartItemHandler.Delete)

	// transaction-product
	transactionProductGroup := r.Group("/transaction-products")
	transactionProductGroup.Use(lib.AuthMiddleware())
	transactionProductGroup.GET("/:transaction_id", transactionProductHandler.GetByTransactionID)
	transactionProductGroup.POST("/:transaction_id", transactionProductHandler.Create)
	transactionProductGroup.DELETE("/:id", transactionProductHandler.Delete)

	//product-variant routes
	productVariantGroup := r.Group("/product-variant")
	productVariantGroup.GET("/:product_id", productVariantHandler.GetByProductID)
	productVariantGroup.GET("/detail/:id", productVariantHandler.GetByID)
	productVariantGroup.Use(lib.AuthMiddleware())
	productVariantGroup.POST("", productVariantHandler.Create)
	productVariantGroup.PATCH("/:id", productVariantHandler.Update)
	productVariantGroup.DELETE("/:id", productVariantHandler.Delete)

	// product size

	productSizeGroup := r.Group("/product-sizes")
	productSizeGroup.GET("/:product_id", productSizeHandler.GetByProductID)
	productSizeGroup.GET("/detail/:id", productSizeHandler.GetByID)
	productSizeGroup.Use(lib.AuthMiddleware())
	productSizeGroup.POST("", productSizeHandler.Create)
	productSizeGroup.PATCH("/:id", productSizeHandler.Update)
	productSizeGroup.DELETE("/:id", productSizeHandler.Delete)

	// product images
	productImagesGroup := r.Group("/product-images")
	productImagesGroup.GET("/:product_id", productImagesHandler.GetByProductID)
	productImagesGroup.Use(lib.AuthMiddleware())
	productImagesGroup.POST("", productImagesHandler.Create)
	productImagesGroup.DELETE("/:id", productImagesHandler.Delete)

	// reviews routes
	reviewsGroup := r.Group("/reviews")
	reviewsGroup.GET("/product/:product_id", reviewsHandler.GetByProductID)
	reviewsGroup.GET("/user/:user_id", reviewsHandler.GetByUserID)
	reviewsGroup.Use(lib.AuthMiddleware())
	reviewsGroup.POST("/:user_id", reviewsHandler.Create)
	reviewsGroup.DELETE("/:id", reviewsHandler.Delete)

}
