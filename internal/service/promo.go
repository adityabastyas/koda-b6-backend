package service

import "koda-b6-backend1/internal/repository"

type PromoService struct {
	repo *repository.PromoRepository
}

func NewPromoService(repo *repository.PromoRepository) *PromoService {
	return &PromoService{repo: repo}
}
