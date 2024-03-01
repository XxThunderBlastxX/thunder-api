package middleware

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func VerifyCaptchaToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUrl := "https://challenges.cloudflare.com/turnstile/v0/siteverify"

		// Get the token from the request header
		token := c.Get("cf-turnstile-response")

		// Data to be sent to Cloudflare API
		req := domain.CFTurnstileToken{
			Secret:   global.Env.CFTurnstileSecret,
			Response: token,
			RemoteIp: c.IP(),
		}

		jsonData, err := json.Marshal(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error marshalling token",
			})
		}

		agent := fiber.Post(reqUrl)
		agent.ContentType("application/json")
		agent.Body(jsonData)

		statusCode, _, resErr := agent.Bytes()
		if resErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while verifying token: ",
			})
		}

		if statusCode != 200 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		return c.Next()
	}
}
