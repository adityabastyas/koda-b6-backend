package repository

import (
	"context"
	"fmt"
	"koda-b6-backend1/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	DB *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
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

func (r *TransactionRepository) GetByID(id int) (*models.Transaction, error) {
	query := `SELECT transaction_id, user_id, promo_id, fullname, email, address, delivery_type, subtotal, tax, total, tanggal FROM transaction WHERE transaction_id = $1`

	rows, err := r.DB.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transaction, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) GetByUserID(userID int) ([]models.Transaction, error) {
	query := `SELECT transaction_id, user_id, promo_id, fullname, email, address, delivery_type, subtotal, tax, total, tanggal FROM transaction WHERE user_id = $1`

	rows, err := r.DB.Query(context.Background(), query, userID)
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

func (r *TransactionRepository) Create(input models.TransactionInput) (*models.Transaction, error) {
	query := `INSERT INTO transaction (user_id, promo_id, fullname, email, address, delivery_type, subtotal, tax, total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING transaction_id, user_id, promo_id, fullname, email, address, delivery_type, subtotal, tax, total, tanggal
	`

	rows, err := r.DB.Query(context.Background(), query, input.UserID, input.PromoID, input.Fullname, input.Email, input.Address, input.DeliveryType, input.Subtotal, input.Tax, input.Total)
	if err != nil {
		fmt.Println("TransactionRepository", err)
		return nil, err
	}
	defer rows.Close()

	transaction, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		fmt.Println("TransactionRepository2", err)
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) Delete(id int) error {
	query := `DELETE FROM transaction WHERE transaction_id = $1`

	_, err := r.DB.Exec(context.Background(), query, id)
	return err
}
