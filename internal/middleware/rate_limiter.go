package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

const (
	localhost      string        = "127.0.0.1"
	maxConnections int           = 5
	expirationTime time.Duration = 1 * time.Minute
)

func RateLimiter() fiber.Handler {
	config := limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == localhost
		},
		Max:        maxConnections,
		Expiration: expirationTime,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{})
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}

	return limiter.New(config)
}
