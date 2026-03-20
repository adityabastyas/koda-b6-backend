package handlers

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(service *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service}
}

// @Summary Ambil semua transaction
// @Tags transaction
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response
// @Router /transactions [get]
func (h *TransactionHandler) GetAll(ctx *gin.Context) {
	trasactions, err := h.service.GetAll()
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
		Result:  trasactions,
	})
}

// @Summary ambil 1 transaction berdasarkan ID
// @Tags transaction
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Response
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "id harus berupa angka",
		})
		return
	}

	transaction, err := h,service.GetByID(id){
		if err != nil{
			ctx.JSON(http.StatusNotFound, models.Response{
				Success: false,
				Message: err.Error(),
			})
		}
	}

	ctx.JSON(http.StatusOK, models.Response{
		success: true,
		Message: "success",
		Result: transaction,
	})
}