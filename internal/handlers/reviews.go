package handlers

import "koda-b6-backend1/internal/service"

type ReviewsHandler struct {
	service service.ReviewsService
}

func NewReviewsHandler(service *service.ReviewsService) *ReviewsHandler {
	return &ReviewsHandler{
		service: *service,
	}
}
