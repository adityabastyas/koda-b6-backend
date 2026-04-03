package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type ProductImagesService struct {
	repo *repository.ProductImagesRepository
}

func NewProductImagesService(repo *repository.ProductImagesRepository) *ProductImagesService {
	return &ProductImagesService{
		repo: repo,
	}
}

func (s *ProductImagesService) GetByProductID(productID int) ([]models.ProductImages, error) {
	if productID <= 0 {
		return nil, errors.New("product id tidak valid")
	}
	return s.repo.GetByProductID(productID)
}
