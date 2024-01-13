package handler

import (
	"github.com/6a6ydoping/online-shop/internal/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/valyala/fasthttp"
)

func (h *Handler) InitRouter(cfg config.RouterConfig) *fasthttp.RequestHandler {
	router := fiber.New()
	router.Use(logger.New(logger.Config{DisableColors: false}))
	router.Get("/", hello)
	r := router.Handler()
	return &r
}

func hello(c fiber.Ctx) error {
	return c.SendString("hello")
}
