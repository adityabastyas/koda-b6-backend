package repository

import (
	"github.com/jackc/pgx/v5"
)

type PromoRepository struct {
	DB *pgx.Conn
}
