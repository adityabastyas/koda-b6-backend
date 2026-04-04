package main

import (
	"fmt"
	"koda-b6-backend1/internal/di"
	"koda-b6-backend1/internal/lib"
	"os"

	_ "koda-b6-backend1/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	godotenv.Load()

	lib.ConnectDB()

	r := gin.Default()

	r.Use(lib.CorsMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	di.Container(r, lib.DB)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
