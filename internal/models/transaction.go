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

// sama persis database
type TransactionProduct struct {
	TransactionProductID int `json:"transaction_product_id" db:"transaction_product_id"`
	TransactionID        int `json:"transaction_id" db:"transaction_id"`
	ProductID            int `json:"product_id" db:"product_id"`
	VariantID            int `json:"variant_id" db:"variant_id"`
	ProductSizeID        int `json:"product_size_id" db:"product_size_id"`
	Quantity             int `json:"quantity" db:"quantity"`
	PriceAtPurchase      int `json:"price_at_purchase" db:"price_at_purchase"`
}

// nerima data dari request body
type TransactionProductInput struct {
	ProductID       int `json:"product_id" binding:"required,gt=0"`
	VariantID       int `json:"variant_id" binding:"required,gt=0"`
	ProductSizeID   int `json:"product_size_id" binding:"required,gt=0"`
	Quantity        int `json:"quantity" binding:"required,gt=0"`
	PriceAtPurchase int `json:"price_at_purchase" binding:"required,gt=0"`
}
