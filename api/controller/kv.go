package controller

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type KVController struct {
	KVService domain.KVService
}

func (kv *KVController) GetValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")

		value, err := kv.KVService.GetValue(key)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   err.Error(),
				Success: false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{
			Data:    &model.SuccessResponse{Message: value},
			Success: true,
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
			return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   err.Error(),
				Success: false,
			})
		}

		if err := kv.KVService.SetKeyValue(req.Key, req.Value); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   err.Error(),
				Success: false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{
			Data:    &model.SuccessResponse{Message: "key set successfully"},
			Success: true,
		})
	}
}

func (kv *KVController) DeleteKeyValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")

		if err := kv.KVService.DeleteKey(key); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{Error: err.Error(), Success: false})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{Data: &model.SuccessResponse{Message: "key deleted successfully"}, Success: true})
	}
}

// TODO: Need to debug this controller
func (kv *KVController) ListKeys() fiber.Handler {
	return func(c *fiber.Ctx) error {
		keys, err := kv.KVService.ListKeys()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{
				Error:   err.Error(),
				Success: false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{
			Data:    &model.SuccessResponse{Message: strings.Join(keys, ", ")},
			Success: true,
		})
	}
}
