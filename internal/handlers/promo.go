package handlers

import "koda-b6-backend1/internal/service"

type PromoHandler struct {
	service *service.PromoService
}

func NewPromoHandler(service *service.PromoService) *PromoHandler {
	return &PromoHandler{service: service}
}
