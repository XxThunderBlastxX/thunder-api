package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	"github.com/XxThunderBlastxX/thunder-api/api/routes"
	appConfig "github.com/XxThunderBlastxX/thunder-api/internal/config"
)

var (
	config *appConfig.AppConfig
)

func init() {
	config = appConfig.NewAppConfig()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: "https://*.koustav.dev"}))
	app.Use(middleware.RateLimiter())
	app.Use(middleware.RequestLogger())

	routes.SetupRoutes(app, config)

	if err := app.Listen(":" + config.AppConfig.Port); err != nil {
		log.Fatal(err)
	}
}
