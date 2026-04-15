package models

type Reviews struct {
	ReviewsID  int     `json:"reviews_id" db:"reviews_id"`
	ProductID  int     `json:"product_id" db:"product_id"`
	UserID     int     `json:"user_id" db:"user_id"`
	Message    string  `json:"message" db:"message"`
	Rating     int     `json:"rating" db:"rating"`
	FullName   string  `json:"full_name" db:"full_name"`
	ProfilePic *string `json:"profile_pic" db:"profile_pic"`
}

type ReviewsInput struct {
	ProductID int    `json:"product_id" binding:"required,gt=0"`
	Message   string `json:"message" binding:"required,min=3"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
}
