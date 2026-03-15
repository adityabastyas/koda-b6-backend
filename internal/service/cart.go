package service

import (
	"errors"
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

func (s *CartService) GetByUserID(userID int) (*models.Cart, error) {
	if userID <= 0 {
		return nil, errors.New("user id tidak valid")
	}
	return s.repo.GetByUserID(userID)
}

func (s *CartService) CreateCart(userID int) error {
	if userID <= 0 {
		return errors.New("user id tidak valid")
	}
	return s.repo.CreateCart(userID)
}
