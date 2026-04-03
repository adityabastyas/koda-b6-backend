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
