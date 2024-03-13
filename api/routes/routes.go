package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	"github.com/XxThunderBlastxX/thunder-api/internal/config"
)

func SetupRoutes(app *fiber.App, config *config.AppConfig) {
	kvBaseURL := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/storage/kv/namespaces/%s", config.AppConfig.Cloudflare.AccountID, config.AppConfig.Cloudflare.KvNamespaceID)

	// Public Routes
	publicRouter := app.Group("/")
	AppRouter(publicRouter, &config.Timer)
	RedirectRouter(publicRouter, kvBaseURL, config.AppConfig.Cloudflare)
	ContactMeRouter(publicRouter, config)

	app.Use(middleware.NewJWTMiddleware(config.AppConfig.Keycloak))

	// Private Routes (Requires Authorization to access these routes)
	privateRouter := app.Group("/")
	KVRouter(privateRouter, kvBaseURL, config.AppConfig.Cloudflare)
}
