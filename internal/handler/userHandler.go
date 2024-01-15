package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) getAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers(c.Context())
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	c.JSON(users)

	return nil
}
