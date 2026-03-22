package repository

import "github.com/jackc/pgx/v5"

type CartItemRepository struct {
	DB *pgx.Conn
}

func NewCartItemRepository(db *pgx.Conn) *CartItemRepository {
	return &CartItemRepository{
		DB: db,
	}
}
