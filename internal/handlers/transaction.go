package handlers

import (
	"errors"
	"fmt"
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
			Message: "gagal mengambil data transaction",
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

	transaction, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "data transaction tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  transaction,
	})
}

// @Summary Ambil semua transaction berdasarkan user ID
// @Tags transaction
// @Produce json
// @Security BearerAuth
// @Param user_id path int true "User ID"
// @Success 200 {object} models.Response
// @Router /transactions/user/{user_id} [get]
func (h *TransactionHandler) GetByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "user id harus berupa angka",
		})
		return
	}

	transactions, err := h.service.GetByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "data transaction tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  transactions,
	})
}

// @Summary Bikin transaction baru
// @Tags transaction
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param transaction body models.TransactionInput true "Transaction Input"
// @Success 200 {object} models.Response
// @Router /transactions [post]
func (h *TransactionHandler) Create(ctx *gin.Context) {
	var input models.TransactionInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Println("error TransactionHandler", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	userID, exist := ctx.Get("user_id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: errors.New("silahkan login terlebih dahulu").Error(),
		})
		return
	}

	input.UserID = userID.(int)

	fmt.Println(input)

	transaction, err := h.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "gagal membuat transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "transaksi berhasil di buat",
		Result: gin.H{
			"transaction_id": transaction.TransactionID,
		},
	})
}

// @Summary Hapus transaction
// @Tags transaction
// @Produce json
// @Security BearerAuth
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.Response
// @Router /transactions/{id} [delete]
func (h *TransactionHandler) Delete(ctx *gin.Context) {
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
			Message: "gagal menghapus transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "transaksi berhasil dihapus",
	})
}
