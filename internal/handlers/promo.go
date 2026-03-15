package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PromoHandler struct {
	service *service.PromoService
}

func NewPromoHandler(service *service.PromoService) *PromoHandler {
	return &PromoHandler{service: service}
}

// @Summary Ambil semua promo
// @Tags promo
// @Produce json
// @Success 200 {object} models.Response
// @Router /promos [get]
func (h *PromoHandler) GetAll(ctx *gin.Context) {
	promos, err := h.service.GetAll()
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
		Result:  promos,
	})
}

// @Summary Ambil 1 promo berdasarkan ID
// @Tags promo
// @Produce json
// @Param id path int true "Promo ID"
// @Success 200 {object} models.Response
// @Router /promos/{id} [get]
func (h *PromoHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	promo, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  promo,
	})
}

// @Summary Tambah promo baru
// @Tags promo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param promo body models.PromoInput true "Promo Input"
// @Success 200 {object} models.Response
// @Router /promos [post]
func (h *PromoHandler) Create(ctx *gin.Context) {
	var input models.PromoInput

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
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "promo berhasil ditambahkan",
	})
}

// @Summary Edit promo
// @Tags promo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Promo ID"
// @Param promo body models.PromoInput true "Promo Input"
// @Success 200 {object} models.Response
// @Router /promos/{id} [patch]
func (h *PromoHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	var input models.PromoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	if err := h.service.Update(id, input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "promo berhasil diupdate",
	})
}

// @Summary Hapus promo
// @Tags promo
// @Produce json
// @Security BearerAuth
// @Param id path int true "Promo ID"
// @Success 200 {object} models.Response
// @Router /promos/{id} [delete]
func (h *PromoHandler) Delete(ctx *gin.Context) {
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
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "promo berhasil dihapus",
	})
}
