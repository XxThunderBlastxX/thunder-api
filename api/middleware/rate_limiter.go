package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"github.com/XxThunderBlastxX/thunder-api/internal/model"
)

func RateLimiter() fiber.Handler {
	config := limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        5,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(model.WebResponse[*model.ErrorResponse]{
				Success: false,
				Error:   "rate limit exceeded",
			})
		},
		LimiterMiddleware: limiter.SlidingWindow{},
	}

	return limiter.New(config)
}
