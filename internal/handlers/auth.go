package handlers

import "koda-b6-backend1/internal/service"

type AuthHandler struct {
	userService   *service.UserService
	forgotService *service.ForgotPasswordService
}

func NewAuthHandler(userService *service.UserService, forgotService *service.ForgotPasswordService) *AuthHandler {
	return &AuthHandler{
		userService:   userService,
		forgotService: forgotService,
	}
}
