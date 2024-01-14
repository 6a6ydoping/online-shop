package handler

import service "github.com/6a6ydoping/online-shop/internal/service/impl"

type Handler struct {
	service service.UserManager
}

func New(um service.UserManager) *Handler {
	return &Handler{
		service: um,
	}
}
