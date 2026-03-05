package di

import (
	"koda-b6-backend1/internal/handlers"
	"koda-b6-backend1/internal/repository"
	"koda-b6-backend1/internal/routes"
	"koda-b6-backend1/internal/service"

	"github.com/gin-gonic/gin"
)

func Container(c *gin.Engine) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes.SetupRoutes(c, userHandler)
}
