package repository

import (
	"context"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
)

type TransactionProductRepository struct {
	DB *pgx.Conn
}

func NewTransactionProductRepository(db *pgx.Conn) *TransactionProductRepository {
	return &TransactionProductRepository{
		DB: db,
	}
}

func (r *TransactionProductRepository) GetByTransactionID(transactionID int) ([]models.TransactionProduct, error) {
	query := `SELECT transaction_product_id, transaction_id, product_id, variant_id, product_size_id, quantity, price_at_purchase FROM transaction_product WHERE transaction_id = $1`

	rows, err := r.DB.Query(context.Background(), query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionProduct])
	if err != nil {
		return nil, err
	}

	return products, nil
}
