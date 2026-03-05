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

func (s *UserService) Register(user models.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email & password tidak boleh kosong")
	}

	if s.repo.FindByEmail(user.Email) != nil {
		return errors.New("email sudah terdaftar")
	}

	s.repo.Save(user)
	return nil
}

func (s *UserService) Login(user models.User) (*models.User, error) {
	x := s.repo.FindByEmail(user.Email)
	if x == nil || x.Password != user.Password {
		return nil, errors.New("email atau password salah")
	}
	return x, nil
}

func (s *UserService) GetAll() []models.User {
	return s.repo.GetAll()
}
