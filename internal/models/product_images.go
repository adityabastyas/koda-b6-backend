package models

type ProductImages struct {
	ProductImagesID int    `json:"product_images_id" db:"product_images_id"`
	ProductID       int    `json:"product_id" db:"product_id"`
	Path            string `json:"path" db:"path"`
}

type ProductImagesInput struct {
	ProductID int    `json:"product_id"`
	Path      string `json:"path"`
}
