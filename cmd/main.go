package main

import (
	"koda-b6-backend1/internal/lib"

	_ "koda-b6-backend1/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	godotenv.Load()

	lib.ConnectDB()
	lib.ConnectConn()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// di.Container(r, lib.DB, lib.Conn)

	r.Run("localhost:8888")
}
