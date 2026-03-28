package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type ProductVariantService struct {
	repo *repository.ProductVariantRepository
}

func NewProductVariantService(repo *repository.ProductVariantRepository) *ProductVariantService {
	return &ProductVariantService{
		repo: repo,
	}
}

func (s *ProductVariantService) GetByProductID(productID int) ([]models.ProductVariant, error) {
	if productID <= 0 {
		return nil, errors.New("product id tidak valid")
	}
	return s.repo.GetByProductID(productID)
}

func (s *ProductVariantService) GetByID(id int) (*models.ProductVariant, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetByID(id)
}
