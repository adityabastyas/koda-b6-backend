package repository

import (
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
