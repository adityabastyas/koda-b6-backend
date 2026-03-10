package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KategoryHandler struct {
	service *service.KategoryService
}

func NewKategoryHandler(service *service.KategoryService) *KategoryHandler {
	return &KategoryHandler{
		service: service,
	}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	kategorys, err := h.service.GetAll()
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
		Result:  kategorys,
	})
}
