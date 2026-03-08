package handlers

import (
	"koda-b6-backend1/internal/lib"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "invalid body", Result: nil})
		return
	}

	if err := h.service.Register(user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Success: false, Message: err.Error(), Result: nil})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Success: true, Message: "register success", Result: nil})

}

func (h *UserHandler) Login(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "invalid body"})
		return
	}

	results, err := h.service.Login(user)
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

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users := h.service.GetAll()
	ctx.JSON(http.StatusOK, models.Response{Success: true, Message: "success", Result: users})
}

func (h *UserHandler) UploadPhoto(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "file tidak ditemukan",
		})
		return
	}

	path := "./uploads/" + file.Filename

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal menyimpan file",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "upload success",
		Result:  path,
	})
}
