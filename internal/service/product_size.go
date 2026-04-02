package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type ProductSizeService struct {
	repo *repository.ProductSizeRepository
}

func NewProductSizeService(repo *repository.ProductSizeRepository) *ProductSizeService {
	return &ProductSizeService{
		repo: repo,
	}
}

func (s *ProductSizeService) GetByProductID(productID int) ([]models.ProductSize, error) {
	if productID <= 0 {
		return nil, errors.New("product id tidak valid")
	}
	return s.repo.GetByProductID(productID)
}

func (s *ProductSizeService) GetByID(id int) (*models.ProductSize, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetByID(id)
}
