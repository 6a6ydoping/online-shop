package handler

import "github.com/6a6ydoping/online-shop/internal/service"

type Handler struct {
	service service.UserService
}

func New(us service.UserService) *Handler {
	return &Handler{
		service: us,
	}
}
