package routes

import (
	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/gofiber/fiber/v2"
)

func AppRouter(router fiber.Router) {
	router.Get("/", controller.AppController())
}
