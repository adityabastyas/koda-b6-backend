package handlers

import (
	"fmt"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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

	if file.Size > 2*1024*1024 {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "file terlalu besar (max 2MB)",
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowedExt[ext] {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "extension tidak diizinkan",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal membaca file",
		})
		return
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal membaca file",
		})
		return
	}

	contentType := http.DetectContentType(buffer)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	if !allowedTypes[contentType] {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "tipe file tidak valid",
		})
		return
	}

	uploadDir := "./uploads"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal membuat folder upload",
		})
		return
	}

	filename := fmt.Sprintf("%d%s", time.Now().Unix(), ext)
	path := filepath.Join(uploadDir, filename)

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal menyimpan file",
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

	err = h.service.UpdateProfilePic(userID.(int), path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal update profile",
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
