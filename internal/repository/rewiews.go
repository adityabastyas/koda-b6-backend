package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReviewsRepository struct {
	DB *pgxpool.Pool
}

func NewReviewsRepository(db *pgxpool.Pool) *ReviewsRepository {
	return &ReviewsRepository{
		DB: db,
	}
}

func (r *ReviewsRepository) GetByProductID(productID int) ([]models.Reviews, error) {
	query := `SELECT reviews_id, product_id, user_id, message, rating FROM reviews WHERE product_id = $1`

	rows, err := r.DB.Query(context.Background(), query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Reviews])
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewsRepository) GetByUserID(userID int) ([]models.Reviews, error) {
	query := `SELECT reviews_id, product_id, user_id, message, rating FROM reviews WHERE user_id = $1`

	rows, err := r.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Reviews])
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewsRepository) Create(userID int, input models.ReviewsInput) error {
	query := `INSERT INTO reviews (product_id, user_id, message, rating) VALUES ($1, $2, $3, $4)`

	_, err := r.DB.Exec(context.Background(), query, input.ProductID, userID, input.Message, input.Rating)
	return err
}
