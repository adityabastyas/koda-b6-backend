package di

import (
	"koda-b6-backend1/internal/handlers"
	"koda-b6-backend1/internal/repository"
	"koda-b6-backend1/internal/routes"
	"koda-b6-backend1/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Container(c *gin.Engine, db *pgxpool.Pool, conn *pgx.Conn) {
	//user
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// forgotpasswor
	forgotRepo := repository.NewForgotPasswordRepository(db)
	forgotService := service.NewForgotPasswordService(userRepo, forgotRepo)

	//auth
	authHandler := handlers.NewAuthHandler(userService, forgotService)

	//product
	productRepo := repository.NewProductRepository(conn)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	//kategory
	kategoryRepo := repository.NewKategoryRepository(db)
	kategoryService := service.NewKategoryService(kategoryRepo)
	kategoryHandler := handlers.NewKategoryHandler(kategoryService)

	// promo conn
	promoRepo := repository.NewPromoRepository(conn)
	promoService := service.NewPromoService(promoRepo)
	promoHandler := handlers.NewPromoHandler(promoService)

	// discount
	discountRepo := repository.NewDiscountRepository(conn)
	discountService := service.NewDiscountService(discountRepo)
	discountHandler := handlers.NewDiscountHandler(discountService)

	//cart
	cartRepo := repository.NewCartRepository(conn)
	cartService := service.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	//transaction
	transactionRepo := repository.NewTransactionRepository(conn)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// cart item
	cartItemRepo := repository.NewCartItemRepository(conn)
	cartItemService := service.NewCartItemService(cartItemRepo, cartRepo)
	cartItemHandler := handlers.NewCartItemHandler(cartItemService)

	routes.SetupRoutes(c, authHandler, userHandler, productHandler, kategoryHandler, promoHandler, discountHandler, cartHandler, transactionHandler, cartItemHandler)
}
