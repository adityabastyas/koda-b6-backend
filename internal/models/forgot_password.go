package models

import "time"

type ForgotPassword struct {
	ID        int        `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	Code      string     `json:"code" db:"code"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordInput struct {
	Code        string `json:"code" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
