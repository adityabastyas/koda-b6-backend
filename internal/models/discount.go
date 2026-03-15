package models

type Discount struct {
	DiscountID   int    `json:"discount_id" db:"discount_id"`
	ProductID    int    `json:"product_id" db:"product_id"`
	FlashSale    bool   `json:"flash_sale" db:"flash_sale"`
	Description  string `json:"description" db:"description"`
	DiscountRate int    `json:"discount_rate" db:"discount_rate"`
}

type DiscountInput struct {
	ProductID    int    `json:"product_id"`
	FlashSale    bool   `json:"flash_sale"`
	Description  string `json:"description"`
	DiscountRate int    `json:"discount_rate"`
}
