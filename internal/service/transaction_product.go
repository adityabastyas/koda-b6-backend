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

func (s *TransactionProductService) Create(transacionID int, input models.TransactionProductInput) error {
	if transacionID <= 0 {
		return errors.New("transaction id tidak valid")
	}
	if input.ProductID <= 0 {
		return errors.New("product id tidak valid")
	}
	if input.Quantity <= 0 {
		return errors.New("quantity tidak valid")
	}
	if input.PriceAtPurchase <= 0 {
		return errors.New("harga tidak valid")
	}
	return s.repo.Create(transacionID, input)
}

func (s *TransactionProductService) Delete(TransactionProductID int) error {
	if TransactionProductID <= 0 {
		return errors.New("transaction product id tidak valid")
	}
	return s.repo.Delete(TransactionProductID)
}
