package handler

import (
	"regexp"

	"github.com/XxThunderBlastxX/thunder-api/internal/models"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

const (
	MaxKeyLength   = 512              // Cloudflare KV key limit
	MaxValueLength = 25 * 1024 * 1024 // 25MB - Cloudflare KV value limit
)

var (
	// Key validation regex: alphanumeric, hyphens, underscores, forward slashes, periods
	keyValidationRegex = regexp.MustCompile(`^[a-zA-Z0-9\-_/.]+$`)
)

type KVHandler struct {
	service *service.KVService
}

func NewKVHandler(service *service.KVService) *KVHandler {
	return &KVHandler{
		service: service,
	}
}

// SetKey creates or updates a key-value pair
func (h *KVHandler) SetKey() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.KeyPair

		// Parse request body
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		// Set key-value pair
		// err := h.service.SetKey(req.Key, req.Value)
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		// }

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
	}
}

// GetValue retrieves a value by key
func (h *KVHandler) GetValue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get key from params
		key := c.Params("key")
		if key == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		// Get value from service
		_, err := h.service.GetValue(key)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

// ListKeys retrieves all keys
func (h *KVHandler) ListKeys() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get keys from service
		_, err := h.service.ListKeys()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}

// DeleteKey deletes a key-value pair
func (h *KVHandler) DeleteKey() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get key from params
		key := c.Params("key")
		if key == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
		}

		// Delete key
		err := h.service.DeleteKey(key)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	}
}
