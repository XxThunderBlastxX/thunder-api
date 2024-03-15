package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/model"
)

type RedirectController struct {
	KVService domain.KVService
}

func (r *RedirectController) Redirect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := c.Params("key")

		value, err := r.KVService.GetValue(key)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(model.WebResponse[*model.ErrorResponse]{Error: err.Error(), Success: false})
		}

		return c.Status(http.StatusOK).Redirect(value)
	}
}
