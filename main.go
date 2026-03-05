package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "koda-b6-backend1/docs"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

// @title User API
// @version 1.0
// @description belajar swagger
// @host localhost:8888
// @BasePath /
type Users struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var ListUser []Users

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")

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

		if token == "" || token != "1234" {
			ctx.JSON(401, Response{
				Success: false,
				Message: "unauthorized",
			})
			return
		}

		ctx.Next()
	}
}

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(corsMiddleware())

	// root
	// @Summary root
	// @Tags Root
	// @Success 200 {object} Response
	// @Router /  [get]
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
		})
	})

	userGroup := r.Group("/users")
	userGroup.Use(authMiddleware())

	userGroup.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
			Result:  ListUser,
		})
	})

	// create user
	// @Summary create user
	// @Tags User
	// @Accept json
	// @Produce json
	// @Param body body Users true "user data"
	// @Success 200 {object} Response
	// @Router /users [post]
	userGroup.POST("", func(ctx *gin.Context) {
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

	// get user by id
	// @Summary get user by id
	// @Tags Users
	// @Param id path string true "user id"
	// @Success 200 {object} Response
	// @Router /users/{id} [get]
	userGroup.GET("/:id", func(ctx *gin.Context) {
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

	// update user
	// @Summary update user
	// @Tags Users
	// @Accept json
	// @Param id path int true "user index"
	// @Param body body Users true "user data"
	// @Success 200 {object} Response
	// @Router /users/{id} [patch]
	userGroup.PATCH("/:id", func(ctx *gin.Context) {
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

	// delete user
	// @Summary delete user
	// @Tags Users
	// @Param id path int true "user index"
	// @Success 200 {object} Response
	// @Router /users/{id} [delete]
	userGroup.DELETE("/:id", func(ctx *gin.Context) {
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
	// @Summary register user
	// @Tags Auth
	// @Accept json
	// @Param body body Users true "register data"
	// @Success 200 {object} Response
	// @Router /register [post]
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

	// login
	// @Summary login user
	// @Tags Auth
	// @Accept json
	// @Param body body Users true "login data"
	// @Success 200 {object} Response
	// @Router /login [post]
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
		fmt.Println(ListUser)
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
