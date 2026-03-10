package service

import (
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type KategoryService struct {
	repo *repository.KategoryRepository
}

func NewKategoryService(repo *repository.KategoryRepository) *KategoryService {
	return &KategoryService{
		repo: repo,
	}
}

func (s *KategoryService) GetAll() ([]models.Kategory, error) {
	return s.repo.GetAll()
}
