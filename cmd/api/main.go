package main

import (
	"github.com/XxThunderBlast/thunder-api/global"
	"github.com/XxThunderBlast/thunder-api/internal/env"
	"github.com/XxThunderBlast/thunder-api/internal/timer"
	"github.com/XxThunderBlast/thunder-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func init() {
	// Initialize the global timer for our application
	timer.InitTimer()

	// Load the environment variables
	if loadEnv, err := env.LoadEnv("."); err != nil {
		log.Fatal(err)
		return
	} else {
		global.Env = loadEnv
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format:        "PID-${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		TimeZone:      "Asia/Kolkata",
		DisableColors: false,
	}))

	routes.Routes(app)

	if err := app.Listen(":" + global.Env.APIPort); err != nil {
		log.Fatal(err)
	}
}
