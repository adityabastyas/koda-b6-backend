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

func (r *CartItemRepository) Create(cartID int, input models.CartItemInput) error {
	query := `INSERT INTO cart_item (cart_id, product_id, variant_id, product_size_id, quantity)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := r.DB.Exec(context.Background(), query, cartID, input.ProductID, input.VariantID, input.ProductSizeID, input.Quantity)
	return err
}

func (r *CartItemRepository) Delete(cartItemID int) error {
	query := `DELETE FROM cart_item WHERE cart_item_id = $1`

	_, err := r.DB.Exec(context.Background(), query, cartItemID)
	return err
}
