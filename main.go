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

	r.POST("/users", func(ctx *gin.Context) {
		data := Users{}

		err := ctx.ShouldBindJSON(&data)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "create user failed",
			})
			return
		}

		for x := range ListUser {
			if ListUser[x].Email == data.Email {
				ctx.JSON(400, Response{
					Success: false,
					Message: "email is ready",
				})
				return
			}
		}

		ListUser = append(ListUser, data)

		ctx.JSON(200, Response{
			Success: true,
			Message: "user created",
		})

	})

	r.Run("localhost:8888")

}
