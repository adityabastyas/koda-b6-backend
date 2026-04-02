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

func (s *ProductSizeService) Create(input models.ProductSizeInput) error {
	if input.ProductID <= 0 {
		return errors.New("product id tidak valid")
	}
	if input.Name == "" {
		return errors.New("nama size tidak boleh kosong")
	}
	return s.repo.Create(input)
}

func (s *ProductSizeService) Update(id int, input models.ProductSizeInput) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	if input.Name == "" {
		return errors.New("nama size tidak boleh kosong")
	}
	return s.repo.Update(id, input)
}
