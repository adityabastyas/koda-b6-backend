package models

import "time"

type Transaction struct {
	TransactionID int       `json:"transaction_id" db:"transaction_id"`
	UserID        int       `json:"user_id" db:"user_id"`
	PromoID       *int      `json:"promo_id" db:"promo_id"`
	Fullname      string    `json:"fullname" db:"fullname"`
	Email         string    `json:"email" db:"email"`
	Address       string    `json:"address" db:"address"`
	DeliveryType  string    `json:"delivery_type" db:"delivery_type"`
	Subtotal      int       `json:"subtotal" db:"subtotal"`
	Tax           int       `json:"tax" db:"tax"`
	Total         int       `json:"total" db:"total"`
	Tanggal       time.Time `json:"tanggal" db:"tanggal"`
}

type TransactionInput struct {
	UserID       int    `json:"user_id"`
	PromoID      *int   `json:"promo_id"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	DeliveryType string `json:"delivery_type"`
	Subtotal     int    `json:"subtotal"`
	Tax          int    `json:"tax"`
	Total        int    `json:"total"`
}
