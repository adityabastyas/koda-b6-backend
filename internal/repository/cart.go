package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type CartRepository struct {
	DB *pgx.Conn
}

func NewCartRepository(db *pgx.Conn) *CartRepository {
	return &CartRepository{
		DB: db}
}

func (r *CartRepository) GetAll() ([]models.Cart, error) {
	query := `SELECT cart_id, user_id FROM cart`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	carts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Cart])
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r *CartRepository) GetByUserID(userID int) (*models.Cart, error) {
	query := `SELECT cart_id, user_id FROM cart WHERE user_id = $1`

	rows, err := r.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cart, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Cart])
	if err != nil {
		return nil, err
	}

	return &cart, nil
}
