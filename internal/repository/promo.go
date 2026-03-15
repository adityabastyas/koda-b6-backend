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

func (r *PromoRepository) Update(id int, input models.PromoInput) error {
	query := `UPDATE promo SET title=$1, description=$2, promo_type=$3, discount_value=$4 WHERE promo_id=$5`

	_, err := r.DB.Exec(context.Background(), query, input.Title, input.Description, input.PromoType, input.DiscountValue, id)
	return err
}

func (r *PromoRepository) Delete(id int) error {
	query := `DELETE FROM promo WHERE promo_id = $1`

	_, err := r.DB.Exec(context.Background(), query, id)
	return err
}

func (r *PromoRepository) GetByID(id int) (*models.Promo, error) {
	query := `SELECT promo_id, title, description, promo_type, discount_value FROM promo WHERE promo_id = $1`

	rows, err := r.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	promo, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Promo])
	if err != nil {
		return nil, err
	}

	return &promo, nil
}
