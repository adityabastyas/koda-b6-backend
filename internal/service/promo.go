package service

import (
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
