package handlers

import (
	"github.com/XxThunderBlast/thunder-api/constants"
	"github.com/XxThunderBlast/thunder-api/types"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetKeyListHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var res types.KVResponse
		path := constants.BaseKVPath + "/keys"

		client := fiber.Get(path)
		client.Set("Authorization", "Bearer "+constants.CFToken)
		client.ContentType("application/json")

		statusCode, body, err := client.Bytes()

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		if statusCode != http.StatusOK {
			return c.Status(statusCode).JSON(fiber.Map{
				"error": string(body),
			})
		}

		if err := json.Unmarshal(body, &res); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	}
}

func PutKVHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var res types.KVResponse
		var req []types.KVRequest

		path := constants.BaseKVPath + "/bulk"

		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		agent := fiber.Put(path)
		agent.Set("Authorization", "Bearer "+constants.CFToken)
		agent.ContentType("application/json")

		reqBody, _ := json.Marshal(req)
		agent.Body(reqBody)

		statusCode, body, err := agent.Bytes()

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		if statusCode != http.StatusOK {
			return c.Status(statusCode).JSON(fiber.Map{
				"error": string(body),
			})
		}

		if err := json.Unmarshal(body, &res); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(http.StatusOK).JSON(res)
	}
}
