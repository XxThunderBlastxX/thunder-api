package controller

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/gofiber/fiber/v2"
)

type ContactMeController struct {
	ContactMeService domain.ContactMeService
}

func (cm *ContactMeController) ContactMe() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var m domain.Message

		if err := c.BodyParser(&m); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request",
			})
		}

		if err := cm.ContactMeService.SendMail(m); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Mail sent successfully",
		})
	}
}
