package service

import "koda-b6-backend1/internal/repository"

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
