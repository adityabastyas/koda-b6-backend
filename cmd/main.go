package main

import (
	"koda-b6-backend1/internal/di"
	"koda-b6-backend1/internal/lib"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("file .env tidak ketemu")
	}

	lib.ConnectDB()

	r := gin.Default()

	di.Container(r)

	r.Run("localhost:8888")
}
