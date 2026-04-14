package handlers

import (
	"fmt"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// @Summary Ambil semua product
// @Tags product
// @Produce json
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /products [get]
func (h *ProductHandler) GetAll(ctx *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "terjadi kesalahan pada server",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  products,
	})
}

// @Summary Ambil 1 product berdasarkan ID
// @Tags product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /products/{id} [get]
func (h *ProductHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "product tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"result":  product,
		"links": gin.H{
			"self":   fmt.Sprintf("/products/%d", product.ProductID),
			"update": fmt.Sprintf("/products/%d", product.ProductID),
			"delete": fmt.Sprintf("/products/%d", product.ProductID),
		},
	})
}

// @Summary Tambah product baru
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body models.ProductInput true "Product Input"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products [post]
func (h *ProductHandler) Create(ctx *gin.Context) {
	var input models.ProductInput

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
			Message: "gagal menambahkan product",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "product berhasil ditambahkan",
	})
}

// @Summary Edit product
// @Tags product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Param product body models.ProductInput true "Product Input"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products/{id} [patch]
func (h *ProductHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	var input models.ProductInput
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
			Message: "gagal update product",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "product berhasil diupdate",
	})
}

// @Summary Hapus product
// @Tags product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products/{id} [delete]
func (h *ProductHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
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
			Message: "gagal menghapus product",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "product berhasil dihapus",
	})
}
