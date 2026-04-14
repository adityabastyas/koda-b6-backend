package models

import "time"

type User struct {
	UserID     int       `json:"user_id" db:"user_id"`
	FullName   string    `json:"full_name" db:"full_name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	Address    *string   `json:"address" db:"address"`
	Phone      *string   `json:"phone" db:"phone"`
	ProfilePic *string   `json:"profile_pic" db:"profile_pic"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Role       string    `json:"role" db:"role"`
}

type UserRegisterInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type UserUpdateInput struct {
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
	Address  *string `json:"address"`
}
