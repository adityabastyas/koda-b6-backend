package handlers

import "koda-b6-backend1/internal/service"

type KategoryHandler struct {
	service *service.KategoryService
}

func NewKategoryHandler(service *service.KategoryService) *KategoryHandler {
	return &KategoryHandler{
		service: service,
	}
}
