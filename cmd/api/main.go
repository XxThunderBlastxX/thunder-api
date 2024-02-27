package main

import (
	"github.com/XxThunderBlast/thunder-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Start the server
	app := fiber.New()

	routes.Routes(app)

	PORT := os.Getenv("PORT")
	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal(err)
	}

}
