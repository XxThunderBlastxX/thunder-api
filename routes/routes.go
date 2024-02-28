package routes

import (
	"github.com/XxThunderBlast/thunder-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.AppHandler())

	app.Put("/kv", handlers.PutKVHandler())
}
