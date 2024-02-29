package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public Routes
	publicRouter := app.Group("/")
	AppRouter(publicRouter)

	// Private Routes (Requires Authorization to access these routes)
	privateRouter := app.Group("/")
	RedirectRouter(privateRouter)
}
