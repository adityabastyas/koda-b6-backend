package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}

	return s.repo.GetByID(id)
}

func (s *ProductService) Create(input models.ProductInput) error {
	if input.Name == "" {
		return errors.New("nama product tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga product tidak valid")
	}

	return s.repo.Create(input)
}

func (s *ProductService) Update(id int, input models.ProductInput) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	if input.Name == "" {
		return errors.New("nama product tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga product tidak valid")
	}

	return s.repo.Update(id, input)
}

func (s *ProductService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	return s.repo.Delete(id)
}
