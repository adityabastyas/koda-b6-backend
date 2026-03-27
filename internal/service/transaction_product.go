package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type TransactionProductService struct {
	repo *repository.TransactionProductRepository
}

func NewTransactionProductService(repo *repository.TransactionProductRepository) *TransactionProductService {
	return &TransactionProductService{
		repo: repo,
	}
}

func (s *TransactionProductService) GetByTransactionID(transactionID int) ([]models.TransactionProduct, error) {
	if transactionID <= 0 {
		return nil, errors.New("transaction id tidak valid")
	}
	return s.repo.GetByTransactionID(transactionID)
}
