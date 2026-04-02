package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductSizeHandler struct {
	service *service.ProductSizeService
}

func NewProductSizeHandler(service *service.ProductSizeService) *ProductSizeHandler {
	return &ProductSizeHandler{
		service: service,
	}
}

// @Summary Ambil semua size berdasarkan product ID
// @Tags product-size
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.Response
// @Router /product-sizes/{product_id} [get]
func (h *ProductSizeHandler) GetByProductID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "product id harus berupa angka",
		})
		return
	}

	sizes, err := h.service.GetByProductID(productID)
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
		Result:  sizes,
	})
}

// @Summary Ambil 1 size berdasarkan ID
// @Tags product-size
// @Produce json
// @Param id path int true "Size ID"
// @Success 200 {object} models.Response
// @Router /product-sizes/detail/{id} [get]
func (h *ProductSizeHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	size, err := h.service.GetByID(id)
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
		Result:  size,
	})
}

// @Summary Tambah size baru
// @Tags product-size
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.ProductSizeInput true "Product Size Input"
// @Success 200 {object} models.Response
// @Router /product-sizes [post]
func (h *ProductSizeHandler) Create(ctx *gin.Context) {
	var input models.ProductSizeInput

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
		Message: "size berhasil ditambahkan",
	})
}

// @Summary Edit size
// @Tags product-size
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Size ID"
// @Param input body models.ProductSizeInput true "Product Size Input"
// @Success 200 {object} models.Response
// @Router /product-sizes/{id} [patch]
func (h *ProductSizeHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	var input models.ProductSizeInput
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
		Message: "size berhasil diupdate",
	})
}

func (h *ProductSizeHandler) Delete(ctx *gin.Context) {
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
		Message: "size berhasil dihapus",
	})
}
