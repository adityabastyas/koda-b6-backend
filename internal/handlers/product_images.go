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
