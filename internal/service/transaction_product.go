package service

import "koda-b6-backend1/internal/repository"

type TransactionProductService struct {
	repo *repository.TransactionProductRepository
}

func NewTransactionProductService(repo *repository.TransactionProductRepository) *TransactionProductService {
	return &TransactionProductService{
		repo: repo,
	}
}
