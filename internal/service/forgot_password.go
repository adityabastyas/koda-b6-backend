package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
	"math/rand"
	"strconv"
)

type ForgotPasswordService struct {
	userRepo   *repository.UserRepository
	forgotRepo *repository.ForgotPasswordRepository
}

func NewForgotPasswordService(userRepo *repository.UserRepository, forgotRepo *repository.ForgotPasswordRepository) *ForgotPasswordService {
	return &ForgotPasswordService{
		userRepo:   userRepo,
		forgotRepo: forgotRepo,
	}
}

func (s *ForgotPasswordService) RequestForgotPassword(input models.ForgotPasswordInput) (string, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil || user == nil {
		return "", errors.New("email tidak ditemukan")
	}

	code := strconv.Itoa(rand.Intn(900000) + 100000)

	if err := s.forgotRepo.CreateForgotRequest(input, code); err != nil {
		return "", errors.New("gagal membuat request forgot password")
	}

	return code, nil
}
