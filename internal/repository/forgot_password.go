package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ForgotPasswordRepository struct {
	DB *pgxpool.Pool
}

func NewForgotPasswordRepository(db *pgxpool.Pool) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{
		DB: db,
	}
}

func (r *ForgotPasswordRepository) GetDataByEmailAndCode(email, code string) (*models.ForgotPassword, error) {
	query := `SELECT id, email, code, created_at, updated_at, deleted_at FROM forgot_password WHERE email = $1 AND code = $2`

	rows, err := r.DB.Query(context.Background(), query, email, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return nil, err
	}

	return &data, nil
}
