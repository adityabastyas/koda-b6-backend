package service

import (
	"errors"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
)

type CartItemService struct {
	cartItemRepo *repository.CartItemRepository
	cartRepo     *repository.CartRepository
}

func NewCartItemService(cartItemRepo *repository.CartItemRepository, cartRepo *repository.CartRepository) *CartItemService {
	return &CartItemService{
		cartItemRepo: cartItemRepo,
		cartRepo:     cartRepo,
	}
}

func (s *CartItemService) GetByUserID(userID int) ([]models.CartItem, error) {
	if userID <= 0 {
		return nil, errors.New("user id tidak valid")
	}

	cart, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, errors.New("cart tidak ditemukan")
	}

	return s.cartItemRepo.GetByCartID(cart.CartID)
}
