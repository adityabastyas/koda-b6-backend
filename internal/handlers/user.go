package handlers

import (
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

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{Success: false, Message: err.Error()})
		return
	}
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

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	var input models.UserUpdateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	userID, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "unauthorized",
		})
		return
	}

	err := h.service.UpdateProfile(userID.(int), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "profile berhasil diupdate",
	})
}
