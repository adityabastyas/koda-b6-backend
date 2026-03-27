package handlers

import "koda-b6-backend1/internal/service"

type TransactionProductHandler struct {
	service *service.TransactionProductService
}

func NewTransactionProductHandler(service *service.TransactionProductService) *TransactionProductHandler {
	return &TransactionProductHandler{
		service: service,
	}
}
