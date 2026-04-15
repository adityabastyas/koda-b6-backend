package models

type Promo struct {
	PromoID       int    `json:"promo_id" db:"promo_id"`
	Title         string `json:"title" db:"title"`
	Description   string `json:"description" db:"description"`
	PromoType     string `json:"promo_type db:"promo_type"`
	DiscountValue int    `json:"discount_value" db:"discount_value"`
}

type PromoInput struct {
	Title         string `json:"title" binding:"required,min=3"`
	Description   string `json:"description" binding:"required"`
	PromoType     string `json:"promo_type" binding:"required,oneof=percentage nominal"`
	DiscountValue int    `json:"discount_value" binding:"required,gt=0"`
}
