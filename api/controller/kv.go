package controller

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/gofiber/fiber/v2"
)

type KVController struct {
	KVService domain.KVService
}

func (kv *KVController) GetValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")

		value, err := kv.KVService.GetValue(key)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"value": value,
		})
	}
}

func (kv *KVController) SetKeyValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		})

		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}

		if err := kv.KVService.SetKeyValue(req.Key, req.Value); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "key-value pair added successfully",
		})
	}
}

func (kv *KVController) DeleteKeyValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")

		if err := kv.KVService.DeleteKey(key); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "key-value pair deleted successfully",
		})
	}
}

// TODO: Need to debug this controller
func (kv *KVController) ListKeys() fiber.Handler {
	return func(c *fiber.Ctx) error {
		keys, err := kv.KVService.ListKeys()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"keys": keys,
		})
	}
}
