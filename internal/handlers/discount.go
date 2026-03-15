package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	service *service.DiscountService
}

func NewDiscountHandler(service *service.DiscountService) *DiscountHandler {
	return &DiscountHandler{service: service}
}

// @Summary Ambil 1 discount berdasarkan ID
// @Tags discount
// @Produce json
// @Param id path int true "Discount ID"
// @Success 200 {object} models.Response
// @Router /discounts/{id} [get]
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
