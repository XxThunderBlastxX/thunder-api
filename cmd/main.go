package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/XxThunderBlastxX/thunder-api/api/routes"
	"github.com/XxThunderBlastxX/thunder-api/internal/db"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/appconfig"
	"github.com/XxThunderBlastxX/thunder-api/internal/global"
	"github.com/XxThunderBlastxX/thunder-api/internal/timer"
)

func init() {
	// Initialize the global timer for our application
	timer.InitTimer()

	// Load the application configuration
	if config, err := appconfig.LoadFromPath(context.TODO(), "internal/config/config.pkl"); err != nil {
		log.Fatal(err)
	} else {
		global.Config = config
	}

	// Connect to MongoDB
	if mongoDb, err := db.ConnectMongo(); err != nil {
		log.Fatal(err)
	} else {
		global.DB = mongoDb
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

	routes.SetupRoutes(app)

	if err := app.Listen(":" + global.Config.Port); err != nil {
		log.Fatal(err)
	}
}
