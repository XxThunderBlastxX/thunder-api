package handler

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type RedirectHandler struct {
	kvService *service.KVService
}

func NewRedirectHandler(service *service.KVService) *RedirectHandler {
	return &RedirectHandler{
		kvService: service,
	}
}

func (h *RedirectHandler) Redirect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the key from the request parameters
		key := c.Params("key")
		if key == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Key is required",
			})
		}

		// Retrieve the value associated with the key
		value, err := h.kvService.GetValue(key)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if value == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Key not found",
			})
		}

		return c.Redirect(value, fiber.StatusPermanentRedirect)
	}
}
