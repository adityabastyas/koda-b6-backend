package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionProductHandler struct {
	service *service.TransactionProductService
}

func NewTransactionProductHandler(service *service.TransactionProductService) *TransactionProductHandler {
	return &TransactionProductHandler{
		service: service,
	}
}

func (h *TransactionProductHandler) GetByTransactionID(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("transaction_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "transaction id harus berupa angka",
		})
		return
	}

	products, err := h.service.GetByTransactionID(transactionID)
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
		Result:  products,
	})

}
