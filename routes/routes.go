package routes

import (
	"github.com/XxThunderBlast/thunder-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", handlers.AppHandler())

	kv := app.Group("/kv")
	kv.Put("/put", handlers.PutKVHandler())
	app.Get("/:key", handlers.RedirectKVHandler())
}
