package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	"github.com/XxThunderBlastxX/thunder-api/api/routes"
	appConfig "github.com/XxThunderBlastxX/thunder-api/internal/config"
)

func main() {
	// Initializing app configuration
	config := appConfig.NewAppConfig()

	// Initializing fiber app
	app := fiber.New()

	// Middlewares
	app.Use(cors.New(cors.Config{}))
	app.Use(middleware.RateLimiter())
	app.Use(middleware.RequestLogger())

	// Setting up routes
	routes.SetupRoutes(app, config)

	// Defer closing the database connection
	defer func(appConfig *appConfig.AppConfig) {
		err := appConfig.Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(config)

	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatal(err)
	}
}
