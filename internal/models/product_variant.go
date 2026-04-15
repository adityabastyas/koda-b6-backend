package models

type ProductVariant struct {
	VariantID   int    `json:"variant_id" db:"variant_id"`
	ProductID   int    `json:"product_id" db:"product_id"`
	Temperature string `json:"temperature" db:"temperature"`
	AddPrice    int    `json:"add_price" db:"add_price"`
}

type ProductVariantInput struct {
	ProductID   int    `json:"product_id" binding:"required,gt=0"`
	Temperature string `json:"temperature" binding:"required,oneof=hot cold ice"`
	AddPrice    int    `json:"add_price" binding:"gte=0"`
}
