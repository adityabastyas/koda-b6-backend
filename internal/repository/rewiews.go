package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ReviewsRepository struct {
	DB *pgxpool.Pool
}

func NewReviewsRepository(db *pgxpool.Pool) *ReviewsRepository {
	return &ReviewsRepository{
		DB: db,
	}
}
