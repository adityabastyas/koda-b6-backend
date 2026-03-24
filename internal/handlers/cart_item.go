package handlers

import "koda-b6-backend1/internal/service"

type CartItemHandler struct {
	service *service.CartItemService
}

func NewCartItemHandler(service *service.CartItemService) *CartItemHandler {
	return &CartItemHandler{
		service: service,
	}
}
