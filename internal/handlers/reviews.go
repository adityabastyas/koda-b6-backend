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

// @Summary Ambil semua review berdasarkan product ID
// @Tags reviews
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.Response
// @Router /reviews/product/{product_id} [get]
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
			Message: "data review tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  reviews,
	})
}

// @Summary Ambil semua review berdasarkan user ID
// @Tags reviews
// @Produce json
// @Security BearerAuth
// @Param user_id path int true "User ID"
// @Success 200 {object} models.Response
// @Router /reviews/user/{user_id} [get]
func (h *ReviewsHandler) GetByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "user id harus berupa angka",
		})
		return
	}

	reviews, err := h.service.GetByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "data review tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  reviews,
	})
}

// @Summary Tambah review baru
// @Tags reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path int true "User ID"
// @Param input body models.ReviewsInput true "Reviews Input"
// @Success 200 {object} models.Response
// @Router /reviews/{user_id} [post]
func (h *ReviewsHandler) Create(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "user id harus berupa angka",
		})
		return
	}

	var input models.ReviewsInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "invalid body",
		})
		return
	}

	if err := h.service.Create(userID, input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "gagal menambahkan review",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "review berhasil ditambahkan",
	})
}

// @Summary Hapus review
// @Tags reviews
// @Produce json
// @Security BearerAuth
// @Param id path int true "Review ID"
// @Success 200 {object} models.Response
// @Router /reviews/{id} [delete]
func (h *ReviewsHandler) Delete(ctx *gin.Context) {
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
			Message: "gagal menghapus review",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "review berhasil dihapus",
	})
}

// @Summary Ambil semua review
// @Tags reviews
// @Produce json
// @Success 200 {object} models.Response
// @Router /reviews [get]
func (h *ReviewsHandler) GetAll(ctx *gin.Context) {
	reviews, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "gagal mengambil data review",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "success",
		Result:  reviews,
	})
}
