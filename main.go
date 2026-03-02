package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type Users struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var ListUser []Users

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}

	}
}

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(401, Response{
				Success: false,
				Message: "unauthorized",
			})
		} else if token != "1234" {
			ctx.JSON(401, Response{
				Success: false,
				Message: "unauthorized",
			})
		} else {
			ctx.Next()
		}

		ctx.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
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

	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "9" {
			ctx.JSON(200, Response{
				Success: true,
				Message: "welcome to the admind",
			})
		} else {
			ctx.JSON(400, Response{
				Success: false,
				Message: fmt.Sprintf("id kamu adalah %s", id),
			})
		}
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		i, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || i < 0 || i >= len(ListUser) {
			ctx.JSON(404, Response{
				Success: false,
				Message: "user not fonud",
			})
			return
		}

		data := Users{}
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "invalid body",
			})
			return
		}

		ListUser[i] = data
		ctx.JSON(200, Response{
			Success: true,
			Message: "update succes",
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		i, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || i < 0 || i >= len(ListUser) {
			ctx.JSON(400, Response{
				Success: false,
				Message: "user not found",
			})
			return
		}

		ListUser = append(ListUser[:i], ListUser[i+1:]...)

		ctx.JSON(200, Response{
			Success: true,
			Message: "user deleted",
		})
	})

	/// register
	r.POST("/register", func(ctx *gin.Context) {
		data := Users{}

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{false, "invalid body", nil})
			return
		}

		if strings.TrimSpace(data.Email) == "" {
			ctx.JSON(400, Response{false, "email wajib diisi", nil})
			return
		}

		if strings.TrimSpace(data.Password) == "" {
			ctx.JSON(400, Response{false, "password wajib diisi", nil})
			return
		}

		for _, u := range ListUser {
			if u.Email == data.Email {
				ctx.JSON(400, Response{false, "email sudah terdaftar", nil})
				return
			}
		}

		ListUser = append(ListUser, data)

		ctx.JSON(200, Response{true, "register success", nil})

	})

	r.POST("/login", func(ctx *gin.Context) {
		data := Users{}

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{false, "invalid body", nil})
			return
		}

		if data.Email == "" || data.Password == "" {
			ctx.JSON(400, Response{false, "email & password wajib", nil})
			return
		}

		for _, u := range ListUser {
			if u.Email == data.Email && u.Password == data.Password {
				ctx.JSON(200, Response{true, "login success", u})
				return
			}
		}

		ctx.JSON(401, Response{false, "email / password salah", nil})
	})

	r.Run("localhost:8888")

}
