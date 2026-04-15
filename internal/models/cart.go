package models

type Cart struct {
	CartID int `json:"cart_id" db:"cart_id"`
	UserID int `json:"user_id" db:"user_id"`
}

type CartItem struct {
	CartItemID    int `json:"cart_item_id" db:"cart_item_id"`
	CartID        int `json:"cart_id" db:"cart_id"`
	ProductID     int `json:"product_id" db:"product_id"`
	VariantID     int `json:"variant_id" db:"variant_id"`
	ProductSizeID int `json:"product_size_id" db:"product_size_id"`
	Quantity      int `json:"quantity" db:"quantity"`
}

type CartItemInput struct {
	ProductID     int `json:"product_id" binding:"required,gt=0"`
	VariantID     int `json:"variant_id" binding:"required,gt=0"`
	ProductSizeID int `json:"product_size_id" binding:"required,gt=0"`
	Quantity      int `json:"quantity" binding:"required,gt=0"`
}
