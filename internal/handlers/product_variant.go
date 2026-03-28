package handlers

import "koda-b6-backend1/internal/service"

type ProductVariantHandler struct {
	service *service.ProductVariantService
}

func NewProductVariantHandler(service *service.ProductVariantService) *ProductVariantHandler {
	return &ProductVariantHandler{
		service: service,
	}
}
