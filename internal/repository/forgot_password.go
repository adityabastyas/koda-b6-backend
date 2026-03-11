package repository

import "github.com/jackc/pgx/v5/pgxpool"

type ForgotPasswordRepository struct {
	DB *pgxpool.Pool
}

func NewForgotPasswordRepository(db *pgxpool.Pool) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{
		DB: db,
	}
}
