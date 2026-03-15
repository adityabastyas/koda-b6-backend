package models

type Cart struct {
	CartID int `json:"cart_id" db:"cart_id"`
	UserID int `json:"user_id" db:"user_id"`
}
