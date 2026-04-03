package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewsHandler struct {
	service service.ReviewsService
}

func NewReviewsHandler(service *service.ReviewsService) *ReviewsHandler {
	return &ReviewsHandler{
		service: *service,
	}
}

func (h *ReviewsHandler) GetByProductID(ctx *gin.Context) {
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "product id harus berupa angka",
		})
		return
	}

	reviews, err := h.service.GetByProductID(productID)
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
		Result:  reviews,
	})
}
