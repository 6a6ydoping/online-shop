package handler

import (
	"github.com/6a6ydoping/online-shop/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) getAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers(c.Context())
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	c.JSON(users)

	return nil
}

func (h *Handler) createUser(c *fiber.Ctx) error {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}
	h.service.CreateUser(c.Context(), *user)

	return nil
}
