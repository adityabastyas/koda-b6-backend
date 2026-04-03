package service

import "koda-b6-backend1/internal/repository"

type ProductImagesService struct {
	repo *repository.ProductImagesRepository
}

func NewProductImagesService(repo *repository.ProductImagesRepository) *ProductImagesService {
	return &ProductImagesService{
		repo: repo,
	}
}
