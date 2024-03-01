package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/timer"
)

func AppController() fiber.Handler {
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
