package service

import (
	"koda-b6-backend1/internal/repository"
)

type DiscountService struct {
	repo *repository.DiscountRepository
}

func NewDiscountService(repo *repository.DiscountRepository) *DiscountService {
	return &DiscountService{
		repo: repo}
}
