package handlers

import (
	"fmt"
	"koda-b6-backend1/internal/lib"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService   *service.UserService
	forgotService *service.ForgotPasswordService
}

func NewAuthHandler(userService *service.UserService, forgotService *service.ForgotPasswordService) *AuthHandler {
	return &AuthHandler{
		userService:   userService,
		forgotService: forgotService,
	}
}

// Register handle POST /register
// @Summary Register user baru
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserRegisterInput true "User Register Input"
// @Success 200 {object} models.Response
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	var input models.UserRegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "invalid body", Result: nil})
		return
	}

	if err := h.service.Register(input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Success: false, Message: err.Error(), Result: nil})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Success: true, Message: "register success", Result: nil})

}

// @Summary Login user
// @Tags auth
// @Accept json
// @Produce json
// @param user body models.UserLoginInput true "User Login Input"
// @Success 200 {object} models.Response
// @Router /auth/login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var input models.UserLoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "invalid body"})
		return
	}

	results, err := h.service.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response{Success: false, Message: err.Error()})
		return
	}

	token, _ := lib.GenerateToken(results.Email)

	ctx.JSON(http.StatusOK, models.Response{Success: true, Message: "login success", Result: gin.H{
		"user":  results,
		"token": token,
	}})
}

// @Summary Request forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.ForgotPasswordInput true "Email input"
// @Success 200 {object} models.Response
// @Router /auth/forgot-password [post]
func (h *AuthHandler) RequestForgotPassword(ctx *gin.Context) {
	var input models.ForgotPasswordInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	code, err := h.forgotService.RequestForgotPassword(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	fmt.Println("OTP CODE", code)

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "code berhasil dikirim",
	})

}
