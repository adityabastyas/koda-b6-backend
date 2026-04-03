package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ProductImagesRepository struct {
	DB *pgxpool.Pool
}

func NewProductImagesRepository(db *pgxpool.Pool) *ProductImagesRepository {
	return &ProductImagesRepository{
		DB: db,
	}
}
