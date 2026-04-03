package handlers

import "koda-b6-backend1/internal/service"

type ProductImagesHandler struct {
	service *service.ProductImagesService
}

func NewProductImagesHandler(service *service.ProductImagesService) *ProductImagesHandler {
	return &ProductImagesHandler{
		service: service,
	}
}
