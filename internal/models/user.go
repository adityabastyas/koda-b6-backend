package models

import "time"

type User struct {
	UserID     int       `json:"user_id" db:"user_id"`
	FullName   string    `json:"full_name" db:"full_name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	Address    *string   `json:"address" db:"address"`
	Phone      *string   `json:"phone" db:"phone"`
	ProfilePic *string   `json:"profile_pic" db:"profile_pic"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Role       string    `json:"role" db:"role"`
}

type UserRegisterInput struct {
	FullName string `json:"full_name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type UserUpdateInput struct {
	FullName string  `json:"full_name" binding:"omitempty,min=3"`
	Email    string  `json:"email" binding:"omitempty,email"`
	Phone    *string `json:"phone" binding:"omitempty"`
	Address  *string `json:"address" binding:"omitempty"`
}
