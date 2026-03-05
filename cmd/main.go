package main

import (
	"koda-b6-backend1/internal/di"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	di.Container(r)

	r.Run("localhost:8888")
}
