package models

type Product struct {
	ProductID   int    `json:"product_id" db:"product_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	KategoryID  int    `json:"kategory_id" db:"kategory_id"`
	ImageURL    string `json:"image_url" db:"image_url"`
}

type ProductInput struct {
	KategoryID  int    `json:"kategory_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageURL    string `json:"image_url"`
}
