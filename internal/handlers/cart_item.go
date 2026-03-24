package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartItemHandler struct {
	service *service.CartItemService
}

func NewCartItemHandler(service *service.CartItemService) *CartItemHandler {
	return &CartItemHandler{
		service: service,
	}
}

// @Summary Ambil semua item di cart berdasarkan user ID
// Tags cart-item
// @Produce json
// @Security BearerAuth
// @Param user_id path int true "User ID"
// @Success 200 {object} models.Response
// @Router /cart-items/{user_id} [get]
func (h *CartItemHandler) GetByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "user id harus berupa angka",
		})
		return
	}

	items, err := h.service.GetByUserID(userID)
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
		Result:  items,
	})
}
