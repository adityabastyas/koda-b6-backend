package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type KategoryRepository struct {
	db *pgxpool.Pool
}

func NewKategoryRepository(db *pgxpool.Pool) *KategoryRepository {
	return &KategoryRepository{
		db: db,
	}
}

func (r *KategoryRepository) GetAll() ([]models.Kategory, error) {

	query := `SELECT kategory_id, name FROM kategory`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	kategorys, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Kategory])
	if err != nil {
		return nil, err
	}

	return kategorys, nil

}

func (r *KategoryRepository) GetByID(id int) (*models.Kategory, error) {
	query := `SELECT kategory_id, name FROM kategory WHERE kategory_id = $1`

	rows, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kategory, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Kategory])
	if err != nil {
		return nil, err
	}

	return &kategory, nil
}
