package handlers

import (
	"github.com/XxThunderBlast/thunder-api/constants"
	"github.com/gofiber/fiber/v2"
	"time"
)

func AppHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"app":     "Thunder-API",
			"version": "1.0.0",
			"uptime": fiber.Map{
				"start": constants.Timer,
				"hr":    time.Since(constants.Timer).Hours(),
				"min":   time.Since(constants.Timer).Minutes(),
				"sec":   time.Since(constants.Timer).Seconds(),
			},
		})
	}
}
