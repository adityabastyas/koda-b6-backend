package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductVariantHandler struct {
	service *service.ProductVariantService
}

func NewProductVariantHandler(service *service.ProductVariantService) *ProductVariantHandler {
	return &ProductVariantHandler{
		service: service,
	}
}

// @Summary Ambil semua variant berdasarkan ID
// @Tags product-variant
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.Response
// @Router /product-variants/{product_id} [get]
func (h *ProductVariantHandler) GetByProductID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "product id harus berupa angka",
		})
		return
	}

	variants, err := h.service.GetByProductID(productID)
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
		Result:  variants,
	})
}

// @Summary Ambil 1 variant berdasarkan ID
// @Tags product-variant
// @Produce json
// @Param id path int true "Variant ID"
// @Success 200 {object} models.Response
// @Router /product-variants/detail/{id} [get]
func (h *ProductVariantHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	variant, err := h.service.GetByID(id)
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
		Result:  variant,
	})
}

// @Summary Tambah variant baru
// @Tags product-variant
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.ProductVariantInput true "Product Variant Input"
// @Success 200 {object} models.Response
// @Router /product-variants [post]
func (h *ProductVariantHandler) Create(ctx *gin.Context) {
	var input models.ProductVariantInput

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
		Message: "variant berhasil di tambahkan",
	})
}

// @Summary Edit variant
// @Tags product-variant
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Variant ID"
// @Param input body models.ProductVariantInput true "Product Variant Input"
// @Success 200 {object} models.Response
// @Router /product-variants/{id} [patch]
func (h *ProductVariantHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	var input models.ProductVariantInput
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
		Message: "variant berhasil diupdate",
	})
}

// @Summary Hapus variant
// @Tags product-variant
// @Produce json
// @Security BearerAuth
// @Param id path int true "Variant ID"
// @Success 200 {object} models.Response
// @Router /product-variants
func (h *ProductVariantHandler) Delete(ctx *gin.Context) {
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
		Message: "variant berhasil dihapus",
	})
}
