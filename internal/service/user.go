package service

import (
	"errors"
	"fmt"
	"koda-b6-backend1/internal/lib"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"

	"github.com/jackc/pgx/v5"
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

	hashedPassword, err := lib.HashPassword(input.Password)
	if err != nil {
		return errors.New("gagal hash password")
	}
	input.Password = hashedPassword

	return s.repo.Save(input)
}

func (s *UserService) Login(input models.UserLoginInput) (*models.User, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("email atau password salah")
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	fmt.Println("User found:", user.Email)
	fmt.Println("Hash from DB:", user.Password)

	valid, err := lib.VerifyPassword(input.Password, user.Password)
	if err != nil {
		fmt.Println("Verify error:", err) 
		return nil, fmt.Errorf("failed to verify password: %w", err)
	}

	if !valid {
		return nil, fmt.Errorf("invalid email or password")
	}

	return user, nil
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}
