package repository

import (
	"context"
	"koda-b6-backend1/internal/lib"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type KategoryRepository struct{}

func NewKategoryRepository() *KategoryRepository {
	return &KategoryRepository{}
}

func (r *KategoryRepository) GetAll() ([]models.Kategory, error) {

	query := `SELECT kategory_id, name FROM kategory`

	rows, err := lib.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	kategory, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Kategory])
	if err != nil {
		return nil, err
	}

	return kategory, nil

}
