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

// @Summary Tambah item ke cart
// @Tags cart-item
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path int tru "User ID"
// @Param input body models.CartItemInput true "Cart Item Input"
// @Success 200 {object} models.Response
// @Router /cart-items/{user_id} [post]
func (h *CartItemHandler) Create(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "user id harus berupa angka",
		})
		return
	}

	var input models.CartItemInput
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
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "item berhasil ditambahkan ke cart",
	})
}

// @Summary Update quantity item di cart
// @Tags cart-item
// @Accept json
// @Produce json
// @Security BearerAuth
// @param ud path int true "Cart Item ID"
// @param quantity query int true "Quantity"
// @Succes 200 {object} models.Response
// @Router /cart-items/{id} [patch]
func (h *CartItemHandler) Update(ctx *gin.Context) {
	cartItemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "cart item id harus berupa angka",
		})
		return
	}

	quantity, err := strconv.Atoi(ctx.Query("quantity"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "quantity harus berupa angka",
		})
		return
	}

	if err := h.service.Update(cartItemID, quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "quantity berhasil dipdate",
	})
}

func (h *CartItemHandler) Delete(ctx *gin.Context) {
	cartItemID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "catt item id harus berupa angka",
		})
		return
	}

	if err := h.service.Delete(cartItemID); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "item berhasil dihapus dari cart",
	})
}
