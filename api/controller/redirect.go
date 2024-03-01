package controller

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
