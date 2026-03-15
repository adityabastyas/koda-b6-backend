package repository

import "github.com/jackc/pgx/v5"

type CartRepository struct {
	DB *pgx.Conn
}

func NewCartRepository(db *pgx.Conn) *CartRepository {
	return &CartRepository{
		DB: db}
}
