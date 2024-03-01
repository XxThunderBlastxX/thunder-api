package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/internal/model"
)

type ContactMeController struct {
	ContactMeService domain.ContactMeService
}

func (cm *ContactMeController) ContactMe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m domain.Message

		if err := c.BodyParser(&m); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse[model.ErrorResponse]{
				Error:   err.Error(),
				Success: false,
			})
		}

		if err := cm.ContactMeService.SendMail(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[model.ErrorResponse]{
				Error:   err.Error(),
				Success: false,
			})
		}

		return c.Status(fiber.StatusOK).JSON(model.WebResponse[*model.SuccessResponse]{
			Data:    &model.SuccessResponse{Message: "Email sent successfully"},
			Success: true,
		})
	}
}
