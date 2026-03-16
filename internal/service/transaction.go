package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo}
}

func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	return s.repo.GetAll()
}

func (s *TransactionService) GetByID(id int) (*models.Transaction, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetByID(id)
}

func (s *TransactionService) GetByUserID(userID int) ([]models.Transaction, error) {
	if userID <= 0 {
		return nil, errors.New("user id tidak valid")
	}
	return s.repo.GetByUserID(userID)
}

func (s *TransactionService) Create(input models.TransactionInput) error {
	if input.UserID <= 0 {
		return errors.New("user id tidak valid")
	}
	if input.Fullname == "" {
		return errors.New("fullname tidak boleh kosong")
	}
	if input.Email == "" {
		return errors.New("email tidak boleh kosong")
	}
	if input.Address == "" {
		return errors.New("address tidak boleh kosong")
	}
	if input.DeliveryType == "" {
		return errors.New("delivery type tidak boleh kosong")
	}
	if input.Total <= 0 {
		return errors.New("total tidak valid")
	}
	return s.repo.Create(input)
}

func (s *TransactionService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.Delete(id)
}
