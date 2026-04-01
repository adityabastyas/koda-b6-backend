package models

type ProductSize struct {
	ProductSizeID int    `json:"product_size_id" db:"product_size_id"`
	ProductID     int    `json:"product_id" db:"product_id"`
	Name          string `json:"name" db:"name"`
	AddPrice      int    `json:"add_price" db:"add_price"`
}

type ProductSizeInput struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	AddPrice  int    `json:"add_price"`
}
