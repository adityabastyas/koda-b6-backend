package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductVariantRepository struct {
	DB *pgxpool.Pool
}

func NewProductVariantRepository(db *pgxpool.Pool) *ProductVariantRepository {
	return &ProductVariantRepository{
		DB: db,
	}
}

func (r *ProductVariantRepository) GetProductID(productID int) ([]models.ProductVariant, error) {
	query := `SELECT variant_id, product_id, temperature, add_price FROM product_variant WHERE product_id = $1`

	rows, err := r.DB.Query(context.Background(), query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	variants, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductVariant])
	if err != nil {
		return nil, err
	}

	return variants, nil
}
