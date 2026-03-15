package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	service *service.DiscountService
}

func NewDiscountHandler(service *service.DiscountService) *DiscountHandler {
	return &DiscountHandler{service: service}
}

// @Summary Ambil semua discount
// @Tags discount
// @Produce json
// @Success 200 {object} models.Response
// @Router /discounts [get]
func (h *DiscountHandler) GetAll(ctx *gin.Context) {
	discounts, err := h.service.GetAll()
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
		Result:  discounts,
	})
}

// @Summary Ambil 1 discount berdasarkan ID
// @Tags discount
// @Produce json
// @Param id path int true "Discount ID"
// @Success 200 {object} models.Response
// @Router /discounts/{id} [get]
func (h *DiscountHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	discount, err := h.service.GetByID(id)
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
		Result:  discount,
	})
}

// @Summary Tambah discount baru
// @Tags discount
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param discount body models.DiscountInput true "Discount Input"
// @Success 200 {object} models.Response
// @Router /discounts [post]
func (h *DiscountHandler) Create(ctx *gin.Context) {
	var input models.DiscountInput

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
		Message: "discount berhasil ditambahkan",
	})
}

// @Summary Edit discount
// @Tags discount
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Discount ID"
// @Param discount body models.DiscountInput true "Discount Input"
// @Success 200 {object} models.Response
// @Router /discounts/{id} [patch]
func (h *DiscountHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	var input models.DiscountInput
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
		Message: "discount berhasil diupdate",
	})
}

// @Summary Hapus discount
// @Tags discount
// @Produce json
// @Security BearerAuth
// @Param id path int true "Discount ID"
// @Success 200 {object} models.Response
// @Router /discounts/{id} [delete]
func (h *DiscountHandler) Delete(ctx *gin.Context) {
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
		Message: "discount berhasil dihapus",
	})
}
