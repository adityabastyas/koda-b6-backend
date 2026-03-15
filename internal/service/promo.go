package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type PromoService struct {
	repo *repository.PromoRepository
}

func NewPromoService(repo *repository.PromoRepository) *PromoService {
	return &PromoService{repo: repo}
}

func (s *PromoService) GetAll() ([]models.Promo, error) {
	return s.repo.GetAll()
}

func (s *PromoService) GetByID(id int) (*models.Promo, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetByID(id)
}

func (s *PromoService) Update(id int, input models.PromoInput) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	if input.Title == "" {
		return errors.New("title tidak boleh kosong")
	}
	if input.DiscountValue <= 0 {
		return errors.New("discount value tidak valid")
	}
	return s.repo.Update(id, input)
}

func (s *PromoService) Create(input models.PromoInput) error {
	if input.Title == "" {
		return errors.New("title tidak boleh kosong")
	}
	if input.PromoType == "" {
		return errors.New("promo type tidak boleh kosong")
	}
	if input.DiscountValue <= 0 {
		return errors.New("discount value tidak valid")
	}
	return s.repo.Create(input)
}

func (s *PromoService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.Delete(id)
}
