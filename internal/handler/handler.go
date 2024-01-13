package handler

import "github.com/6a6ydoping/online-shop/internal/service"

type Handler struct {
	service service.Service
}

func New(s service.Service) *Handler {
	return &Handler{
		service: s,
	}
}
