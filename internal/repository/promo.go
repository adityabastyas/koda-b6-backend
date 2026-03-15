package repository

import (
	"github.com/jackc/pgx/v5"
)

type PromoRepository struct {
	DB *pgx.Conn
}

func NewPromoRepository(db *pgx.Conn) *PromoRepository {
	return &PromoRepository{
		DB: db,
	}
}
