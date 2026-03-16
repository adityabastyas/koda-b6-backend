package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	DB *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) GetAll() ([]models.Transaction, error) {
	query := `SELECT transaction_id, user_id, promo_id, fullname, email, address, delivery_type, subtotal, tax, total, tanggal FROM transaction`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
