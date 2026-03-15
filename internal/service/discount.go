package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type DiscountService struct {
	repo *repository.DiscountRepository
}

func NewDiscountService(repo *repository.DiscountRepository) *DiscountService {
	return &DiscountService{
		repo: repo}
}

func (s *DiscountService) GetAll() ([]models.Discount, error) {
	return s.repo.GetAll()
}

func (s *DiscountService) GetByID(id int) (*models.Discount, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetByID(id)
}
