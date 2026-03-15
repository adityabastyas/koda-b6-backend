package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type DiscountRepository struct {
	DB *pgx.Conn
}

func NewDiscountRepository(db *pgx.Conn) *DiscountRepository {
	return &DiscountRepository{
		DB: db,
	}
}

func (r *DiscountRepository) GetAll() ([]models.Discount, error) {
	query := `SELECT discount_id, product_id, flash_sale, description, discount_rate FROM discount`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Discount])
	if err != nil {
		return nil, err
	}

	return discounts, nil
}

func (r *DiscountRepository) GetByID(id int) (*models.Discount, error) {
	query := `SELECT discount_id, product_id, flash_sale, description, discount_rate FROM discount WHERE discount_id = $1`

	rows, err := r.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	discount, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Discount])
	if err != nil {
		return nil, err
	}

	return &discount, nil
}

func (r *DiscountRepository) Create(input models.DiscountInput) error {
	query := `INSERT INTO discount (product_id, flash_sale, description, discount_rate) VALUES ($1, $2, $3, $4)`

	_, err := r.DB.Exec(context.Background(), query, input.ProductID, input.FlashSale, input.Description, input.DiscountRate)
	return err
}

func (r *DiscountRepository) Update(id int, input models.DiscountInput) error {
	query := `UPDATE discount SET product_id=$1, flash_sale=$2, description=$3, discount_rate=$4 WHERE discount_id=$5`

	_, err := r.DB.Exec(context.Background(), query, input.ProductID, input.FlashSale, input.Description, input.DiscountRate, id)
	return err
}

func (r *DiscountRepository) Delete(id int) error {
	query := `DELETE FROM discount WHERE discount_id = $1`

	_, err := r.DB.Exec(context.Background(), query, id)
	return err
}
