package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductImagesHandler struct {
	service *service.ProductImagesService
}

func NewProductImagesHandler(service *service.ProductImagesService) *ProductImagesHandler {
	return &ProductImagesHandler{
		service: service,
	}
}

// @Summary Ambil semua image berdasarkan product ID
// @Tags product-images
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.Response
// @Router /product-images/{product_id} [get]
func (h *ProductImagesHandler) GetByProductID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "product id harus berupa angka",
		})
		return
	}

	images, err := h.service.GetByProductID(productID)
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
		Result:  images,
	})
}

// @Summary Tambah image baru
// @Tags product-images
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param input body models.ProductImagesInput true "Product Images Input"
// @Success 200 {object} models.Response
// @Router /product-images [post]
func (h *ProductImagesHandler) Create(ctx *gin.Context) {
	var input models.ProductImagesInput

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
		Message: "image berhasil ditambahkan",
	})
}

// @Summary Hapus image
// @Tags product-images
// @Produce json
// @Security BearerAuth
// @Param id path int true "Image ID"
// @Success 200 {object} models.Response
// @Router /product-images/{id} [delete]
func (h *ProductImagesHandler) Delete(ctx *gin.Context) {
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
		Message: "image berhasil dihapus",
	})
}
