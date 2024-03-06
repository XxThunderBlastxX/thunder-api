package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RequestLogger() fiber.Handler {
	config := logger.Config{
		Format:        "PID-${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		TimeZone:      "Asia/Kolkata",
		DisableColors: false,
	}

	return logger.New(config)
}
