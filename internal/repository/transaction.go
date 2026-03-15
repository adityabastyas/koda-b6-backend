package repository

import "github.com/jackc/pgx/v5"

type TransactionRepository struct {
	DB *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{DB: db}
}
