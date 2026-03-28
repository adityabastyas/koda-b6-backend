package models

type ProductVariant struct {
	VariantID   int    `json:"variant_id" db:"variant_id"`
	ProductID   int    `json:"product_id" db:"product_id"`
	Temperature string `json:"temperature" db:"temperature"`
	AddPrice    int    `json:"add_price" db:"add_price"`
}

type ProductVariantInput struct {
	ProductID   int    `json:"product_id"`
	Temperature string `json:"temperature"`
	AddPrice    int    `json:"add_price"`
}
