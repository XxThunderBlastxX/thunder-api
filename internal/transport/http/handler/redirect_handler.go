package handler

import "github.com/gofiber/fiber/v2"

type RedirectHandler struct{}

func NewRedirectHandler() *RedirectHandler {
	return &RedirectHandler{}
}

func (h *RedirectHandler) Redirect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		panic("Redirect handler not implemented")
	}
}
