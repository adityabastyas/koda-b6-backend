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

func (r *ForgotPasswordRepository) DeleteByCode(code string) error {
	query := `DELETE FROM forgot_password WHERE code = $1`

	_, err := r.DB.Exec(context.Background(), query, code)
	return err

}

func (r *ForgotPasswordRepository) CreateForgotRequest(input models.ForgotPasswordInput, code string) error {
	query := `INSERT INTO forgot_password (email, code) VALUES ($1, $2)`

	_, err := r.DB.Exec(context.Background(), query, input.Email, code)
	return err
}
