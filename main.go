package main

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type Users struct {
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"password"`
}

var ListUser []Users

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "Datas user",
			Result:  ListUser,
		})
	})

	r.Run("localhost:8888")

}
