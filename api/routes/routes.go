package routes

import (
	"github.com/gofiber/fiber/v2"
)

/* TODO : Need to fix this variable
 * Error: panic: runtime error: invalid memory address or nil pointer dereference
 */
var (
	//kvBaseURL = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/storage/kv/namespaces/%s", global.Config.Cloudflare.AccountID, global.Config.Cloudflare.KvNamespaceID)
	kvBaseURL = ""
)

func SetupRoutes(app *fiber.App) {
	// Public Routes
	publicRouter := app.Group("/")
	AppRouter(publicRouter)
	RedirectRouter(publicRouter)

	// Private Routes (Requires Authorization to access these routes)
	privateRouter := app.Group("/")
	KVRouter(privateRouter)
	ContactMeRouter(privateRouter)
}
