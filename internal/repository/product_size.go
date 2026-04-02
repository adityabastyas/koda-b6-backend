package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductSizeRepository struct {
	DB *pgxpool.Pool
}

func NewProductSizeRepository(db *pgxpool.Pool) *ProductSizeRepository {
	return &ProductSizeRepository{
		DB: db,
	}
}

func (r *ProductSizeRepository) GetByProductID(productID int) ([]models.ProductSize, error) {
	query := `SELECT product_size_id, product_id, name, add_price FROM product_size WHERE product_id = $1`

	rows, err := r.DB.Query(context.Background(), query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sizes, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductSize])
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (r *ProductSizeRepository) GetByID(id int) (*models.ProductSize, error) {
	query := `SELECT product_size_id, product_id, name, add_price FROM product_size WHERE product_size_id = $1`

	rows, err := r.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	size, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductSize])
	if err != nil {
		return nil, err
	}

	return &size, nil
}

func (r *ProductSizeRepository) Create(input models.ProductSizeInput) error {
	query := `INSERT INTO product_size (product_id, name, add_price) VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(context.Background(), query, input.ProductID, input.Name, input.AddPrice)
	return err
}
