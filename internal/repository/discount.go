package repository

import "github.com/jackc/pgx/v5"

type DiscountRepository struct {
	DB *pgx.Conn
}

func NewDiscountRepository(db *pgx.Conn) *DiscountRepository {
	return &DiscountRepository{
		DB: db,
	}
}
