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
	ProductID int    `json:"product_id"`
	Message   string `json:"message"`
	Rating    int    `json:"rating"`
}
