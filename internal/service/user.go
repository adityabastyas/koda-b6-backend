package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(input models.UserRegisterInput) error {
	if input.Email == "" || input.Password == "" {
		return errors.New("email & password tidak boleh kosong")
	}

	user, _ := s.repo.FindByEmail(input.Email)
	if user != nil {
		return errors.New("email sudah terdaftar")
	}

	return s.repo.Save(input)
	return nil
}

func (s *UserService) Login(input models.UserLoginInput) (*models.User, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil || user.Password != input.Password {
		return nil, errors.New("email atau password salah")
	}
	return user, nil
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}
