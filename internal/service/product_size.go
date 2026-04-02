package service

import "koda-b6-backend1/internal/repository"

type ProductSizeService struct {
	repo *repository.ProductSizeRepository
}

func NewProductSizeService(repo *repository.ProductSizeRepository) *ProductSizeService {
	return &ProductSizeService{
		repo: repo,
	}
}
