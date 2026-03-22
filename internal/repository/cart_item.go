package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type CartItemRepository struct {
	DB *pgx.Conn
}

func NewCartItemRepository(db *pgx.Conn) *CartItemRepository {
	return &CartItemRepository{
		DB: db,
	}
}

func (r *CartItemRepository) GetByCartID(cartID int) ([]models.CartItem, error) {
	query := `SELECT cart_item_id, cart_id, product_id, variant_id, product_size_id, quantity FROM cart_item WHERE cart_id = $1`

	rows, err := r.DB.Query(context.Background(), query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.CartItem])
	if err != nil {
		return nil, err
	}

	return items, nil
}
