package main

import (
	"fmt"
	"github.com/XxThunderBlast/thunder-api/constants"
	"github.com/XxThunderBlast/thunder-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	constants.Timer = time.Now()

	constants.CFToken = os.Getenv("CF_TOKEN")
	constants.KVNamespaceID = os.Getenv("KV_NAMESPACE_ID")
	constants.CFAccountID = os.Getenv("CF_ID")

	constants.BaseKVPath = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/storage/kv/namespaces/%s", constants.CFAccountID, constants.KVNamespaceID)
}

func main() {

	// Start the server
	app := fiber.New()

	routes.Routes(app)

	PORT := os.Getenv("PORT")
	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal(err)
	}

}
