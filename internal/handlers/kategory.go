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

// @Summary Ambil semua kategory
// @Tags kategory
// @Produce json
// @Success 200 {object} models.Response
// @Router /kategorys [get]
func (h *KategoryHandler) GetAll(ctx *gin.Context) {
	kategorys, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal mengambil data kategory",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  kategorys,
	})
}

// @Summary Ambil kategory berdasarkan ID
// @Tags kategory
// @Produce json
// @Param id path int true "Kategory ID"
// @Success 200 {object} models.Response
// @Router /kategorys/{id} [get]
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
			Message: "data kategory tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  kategory,
	})
}

// @Summary Tambah kategory baru
// @Tags kategory
// @Accept json
// @Produce json
// @Param input body models.KategoryInput true "Kategory Input"
// @Success 200 {object} models.Response
// @Router /kategorys [post]
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
			Message: "gagal menambahkan kategory",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "kategory berhasil ditambahkan",
	})
}

// @Summary Update kategory
// @Tags kategory
// @Accept json
// @Produce json
// @Param id path int true "Kategory ID"
// @Param input body models.KategoryInput true "Kategory Input"
// @Success 200 {object} models.Response
// @Router /kategorys/{id} [patch]
func (h *KategoryHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	var input models.KategoryInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
	}

	if err := h.service.Update(id, input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "gagal update kategory",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "kategory berhasil diupdate",
	})
}

// @Summary Hapus kategory
// @Tags kategory
// @Produce json
// @Param id path int true "Kategory ID"
// @Success 200 {object} models.Response
// @Router /kategorys/{id} [delete]
func (h *KategoryHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "gagal menghapus kategory",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "kategory berhasil dihapus",
	})
}
