package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ProductVariantRepository struct {
	DB *pgxpool.Pool
}

func NewProductVariantRepository(db *pgxpool.Pool) *ProductVariantRepository {
	return &ProductVariantRepository{
		DB: db,
	}
}
