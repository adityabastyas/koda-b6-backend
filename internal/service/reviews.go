package service

import "koda-b6-backend1/internal/repository"

type ReviewsService struct {
	repo *repository.ReviewsRepository
}

func NewReviewsService(repo *repository.ReviewsRepository) *ReviewsService {
	return &ReviewsService{
		repo: repo,
	}
}
