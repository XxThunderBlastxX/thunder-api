package handlers

import (
	"github.com/XxThunderBlast/thunder-api/global"
	"github.com/XxThunderBlast/thunder-api/internal/timer"
	"github.com/gofiber/fiber/v2"
)

func AppHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"app":     "Thunder-API",
			"version": "1.0.0",
			"uptime": fiber.Map{
				"start": global.Timer,
				"hr":    timer.TimeElapsedHR(),
				"min":   timer.TimeElapsedMin(),
				"sec":   timer.TimeElapsedSec(),
			},
		})
	}
}
