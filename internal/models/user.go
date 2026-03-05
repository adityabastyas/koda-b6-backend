package models

// @title User API
// @version 1.0
// @description belajar swagger
// @host localhost:8888
// @BasePath /
type Users struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}
