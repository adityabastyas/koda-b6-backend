package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KategoryHandler struct {
	service *service.KategoryService
}

func NewKategoryHandler(service *service.KategoryService) *KategoryHandler {
	return &KategoryHandler{
		service: service,
	}
}

func (h *KategoryHandler) GetAll(ctx *gin.Context) {
	kategorys, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  kategorys,
	})
}

func (h *KategoryHandler) GetByID(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	kategory, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  kategory,
	})
}

func (h *KategoryHandler) Create(ctx *gin.Context) {
	var input models.KategoryInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	if err := h.service.Create(input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "kategory berhasil ditambahkan",
	})
}
