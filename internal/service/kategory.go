package service

import (
	"errors"
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

func (s *KategoryService) GetByID(id int) (*models.Kategory, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}

	return s.repo.GetByID(id)
}

func (s *KategoryService) Create(input models.KategoryInput) error {
	if input.Name == "" {
		return errors.New("nama kategory tidak boleh kosong")
	}

	return s.repo.Create(input)
}

func (s *KategoryService) Update(id int, input models.KategoryInput) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	if input.Name == "" {
		return errors.New("nama kategory tidak boleh kosong")
	}

	return s.Update(id, input)
}
