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
