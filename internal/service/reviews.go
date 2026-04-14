package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type ReviewsService struct {
	repo *repository.ReviewsRepository
}

func NewReviewsService(repo *repository.ReviewsRepository) *ReviewsService {
	return &ReviewsService{
		repo: repo,
	}
}

func (s *ReviewsService) GetByProductID(productID int) ([]models.Reviews, error) {
	if productID <= 0 {
		return nil, errors.New("product id tidak valid")
	}
	return s.repo.GetByProductID(productID)
}

func (s *ReviewsService) GetByUserID(userID int) ([]models.Reviews, error) {
	if userID <= 0 {
		return nil, errors.New("user id tidak valid")
	}
	return s.repo.GetByUserID(userID)
}

func (s *ReviewsService) Create(userID int, input models.ReviewsInput) error {
	if userID <= 0 {
		return errors.New("user id tidak valid")
	}
	if input.ProductID <= 0 {
		return errors.New("product id tidak valid")
	}
	if input.Message == "" {
		return errors.New("message tidak boleh kosong")
	}

	if input.Rating < 1 || input.Rating > 5 {
		return errors.New("rating harus antara 1-5")
	}
	return s.repo.Create(userID, input)
}

func (s *ReviewsService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.Delete(id)
}

func (s *ReviewsService) GetAll() ([]models.Reviews, error) {
	return s.repo.GetAll()
}
