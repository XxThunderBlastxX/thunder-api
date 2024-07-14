package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"log"

	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	"github.com/XxThunderBlastxX/thunder-api/api/routes"
	appConfig "github.com/XxThunderBlastxX/thunder-api/internal/config"
)

func main() {
	// Initializing app configuration
	config := appConfig.NewAppConfig()

	// Initializing fiber app
	app := fiber.New()

	// Cors Middlewares
	app.Use(cors.New(cors.Config{}))

	// Rate Limiter Middleware
	app.Use(middleware.RateLimiter())

	// Request Logger Middleware
	app.Use(middleware.RequestLogger())

	// Favicon Middleware
	app.Use(favicon.New(favicon.Config{
		Data:         config.Favicon,
		CacheControl: "public, max-age=31536000",
	}))

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
