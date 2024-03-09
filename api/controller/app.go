package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func AppController(appTimer *time.Time) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"app":     "Thunder-API",
			"version": "1.0.0",
			"uptime": fiber.Map{
				"start": appTimer,
				"hr":    time.Since(*appTimer).Hours(),
				"min":   time.Since(*appTimer).Minutes(),
				"sec":   time.Since(*appTimer).Seconds(),
			},
		})
	}
}
