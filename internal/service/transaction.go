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
