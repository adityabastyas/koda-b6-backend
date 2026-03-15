package service

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) GetAll() ([]models.Cart, error) {
	return s.repo.GetAll()
}
