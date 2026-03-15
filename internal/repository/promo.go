package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

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

func (r *PromoRepository) GetAll() ([]models.Promo, error) {
	query := `SELECT promo_id, title, description, promo_type, discount_value FROM promo`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	promos, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Promo])
	if err != nil {
		return nil, err
	}

	return promos, nil
}

func (r *PromoRepository) Create(input models.PromoInput) error {
	query := `INSERT INTO promo (title, description, promo_type, discount_value) VALUES ($1, $2, $3, $4)`

	_, err := r.DB.Exec(context.Background(), query, input.Title, input.Description, input.PromoType, input.DiscountValue)
	return err
}
