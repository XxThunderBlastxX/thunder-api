package middleware

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/model"
)

func VerifyCaptchaToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUrl := "https://challenges.cloudflare.com/turnstile/v0/siteverify"

		// Get the token from the request header
		token := c.Get("cf-turnstile-response")

		// Data to be sent to Cloudflare API
		req := model.CFTurnstileToken{
			Secret:   global.Config.Cloudflare.TurnstileSecret,
			Response: token,
			RemoteIp: c.IP(),
		}

		jsonData, err := json.Marshal(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   "Error while marshaling token: " + err.Error(),
				Success: false,
			})
		}

		agent := fiber.Post(reqUrl)
		agent.ContentType("application/json")
		agent.Body(jsonData)

		statusCode, _, resErr := agent.Bytes()
		if resErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   "Error while sending request to Cloudflare: " + resErr[0].Error(),
				Success: false,
			})
		}

		if statusCode != 200 {
			return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   "Invalid token",
				Success: false,
			})
		}

		return c.Next()
	}
}
