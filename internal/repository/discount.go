package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type DiscountRepository struct {
	DB *pgx.Conn
}

func NewDiscountRepository(db *pgx.Conn) *DiscountRepository {
	return &DiscountRepository{
		DB: db,
	}
}

func (r *DiscountRepository) GetAll() ([]models.Discount, error) {
	query := `SELECT discount_id, product_id, flash_sale, description, discount_rate FROM discount`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Discount])
	if err != nil {
		return nil, err
	}

	return discounts, nil
}

func (r *DiscountRepository) GetByID(id int) (*models.Discount, error) {
	query := `SELECT discount_id, product_id, flash_sale, description, discount_rate FROM discount WHERE discount_id = $1`

	rows, err := r.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discount, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Discount])
	if err != nil {
		return nil, err
	}

	return &discount, nil
}
