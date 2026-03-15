package handlers

import "koda-b6-backend1/internal/service"

type DiscountHandler struct {
	service *service.DiscountService
}

func NewDiscountHandler(service *service.DiscountService) *DiscountHandler {
	return &DiscountHandler{service: service}
}
