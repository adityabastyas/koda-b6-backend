package repository

import "github.com/jackc/pgx/v5"

type TransactionProductRepository struct {
	DB *pgx.Conn
}

func NewTransactionProductRepository(db *pgx.Conn) *TransactionProductRepository {
	return &TransactionProductRepository{
		DB: db,
	}
}
