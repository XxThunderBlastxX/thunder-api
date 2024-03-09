package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
)

func AppRouter(router fiber.Router, timer *time.Time) {
	router.Get("/", controller.AppController(timer))
}
