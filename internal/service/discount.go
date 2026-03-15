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

func (s *DiscountService) Create(input models.DiscountInput) error {
	if input.ProductID <= 0 {
		return errors.New("product id tidak valid")
	}
	if input.DiscountRate <= 0 {
		return errors.New("discount rate tidak valid")
	}
	return s.repo.Create(input)
}

func (s *DiscountService) Update(id int, input models.DiscountInput) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	if input.ProductID <= 0 {
		return errors.New("product id tidak valid")
	}
	if input.DiscountRate <= 0 {
		return errors.New("discount rate tidak valid")
	}
	return s.repo.Update(id, input)
}

func (s *DiscountService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.Delete(id)
}
