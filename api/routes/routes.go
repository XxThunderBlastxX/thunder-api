package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/internal/global"
)

var (
	kvBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%v/storage/kv/namespaces/%v", global.Env.CFAccountId, global.Env.KvNamespaceId)
)

func SetupRoutes(app *fiber.App) {
	// Public Routes
	publicRouter := app.Group("/")
	AppRouter(publicRouter)
	RedirectRouter(publicRouter)

	//Private Routes (Requires Authorization to access these routes)
	privateRouter := app.Group("/")
	KVRouter(privateRouter)
	ContactMeRouter(privateRouter)
}
