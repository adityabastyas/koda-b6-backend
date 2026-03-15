package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"

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
