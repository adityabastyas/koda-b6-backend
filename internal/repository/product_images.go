package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductImagesRepository struct {
	DB *pgxpool.Pool
}

func NewProductImagesRepository(db *pgxpool.Pool) *ProductImagesRepository {
	return &ProductImagesRepository{
		DB: db,
	}
}

func (r *ProductImagesRepository) GetByProductID(productID int) ([]models.ProductImages, error) {
	query := `SELECT product_images_id, product_id, path FROM product_images WHERE product_id = $1`

	rows, err := r.DB.Query(context.Background(), query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	images, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductImages])
	if err != nil {
		return nil, err
	}

	return images, nil
}
