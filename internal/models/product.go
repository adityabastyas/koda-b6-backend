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
	KategoryID  int    `json:"kategory_id" binding:"required,gt=0"`
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=5"`
	Price       int    `json:"price" binding:"required,gt=0"`
	ImageURL    string `json:"image_url" binding:"required,url"`
}
