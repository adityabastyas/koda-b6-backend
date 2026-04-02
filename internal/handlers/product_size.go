package handlers

import "koda-b6-backend1/internal/service"

type ProductSizeHandler struct {
	service *service.ProductSizeService
}

func NewProductSizeHandler(service *service.ProductSizeService) *ProductSizeHandler {
	return &ProductSizeHandler{
		service: service,
	}
}
