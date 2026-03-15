package handlers

import "koda-b6-backend1/internal/service"

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}
