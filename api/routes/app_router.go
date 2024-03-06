package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
)

func AppRouter(router fiber.Router) {
	router.Get("/", controller.AppController())
}
