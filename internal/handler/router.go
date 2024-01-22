package handler

import (
	"github.com/6a6ydoping/online-shop/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp"
)

func (h *Handler) InitRouter(cfg config.RouterConfig) *fasthttp.RequestHandler {
	router := fiber.New()
	router.Use(logger.New(logger.Config{DisableColors: false}))

	router.Get("/allUsers", h.getAllUsers)
	router.Post("user", h.createUser)
	
	r := router.Handler()
	return &r
}
