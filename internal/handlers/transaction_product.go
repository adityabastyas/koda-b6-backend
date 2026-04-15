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

// @Summary Ambil semua product berdasarkan transaction ID
// @Tags transaction-product
// @Produce json
// @Security BearerAuth
// @Param transaction_id path int true "Transaction ID"
// @Success 200 {object} models.Response
// @Router /transaction-products/{transaction_id} [get]
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
			Message: "data product tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  products,
	})

}

// @Summary Tambah product ke transaksi
// @Tags transaction-product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param transaction_id path int true "Transaction ID"
// @Param input body models.TransactionProductInput true "Transaction Product Input"
// @Success 200 {object} models.Response
// @Router /transaction-product/{transaction_id} [post]
func (h *TransactionProductHandler) Create(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("transaction_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "transaction id harus berupa angka",
		})
		return
	}

	var input models.TransactionProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	if err := h.service.Create(transactionID, input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "gagal menambahkan product ke transaksi",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "product berhasil ditambahkab ke transaksi",
	})
}

// @Summary Hapus product dari transaksi
// @Tags transaction-product
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction Product ID"
// @Success 200 {object} models.Response
// @Router /transaction-products/{id} [delete]
func (h *TransactionProductHandler) Delete(ctx *gin.Context) {
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
			Message: "gagal menghapus product dari transaksi",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "product berhasil dihapus dari transaksi",
	})
}
