package handler

import (
	"github.com/gofiber/fiber/v2"
)

type AppHandler struct{}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

func (h *AppHandler) HandleRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res := fiber.Map{
			"message": "Welcome to Thunder API",
			"status":  "success",
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}
}
