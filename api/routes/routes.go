package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	"github.com/XxThunderBlastxX/thunder-api/internal/global"
	"github.com/XxThunderBlastxX/thunder-api/internal/repository"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func SetupRoutes(app *fiber.App) {
	kvBaseURL := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/storage/kv/namespaces/%s", global.Config.Cloudflare.AccountID, global.Config.Cloudflare.KvNamespaceID)

	// Public Routes
	publicRouter := app.Group("/")
	AppRouter(publicRouter)
	RedirectRouter(publicRouter, kvBaseURL)

	keycloakRepo := repository.NewKeycloakRepository(global.Config.Keycloak.AuthUrl, global.Config.Keycloak.Realm, global.Config.Keycloak.ClientId, global.Config.Keycloak.ClientSecret)
	keycloakService := service.NewKeycloakService(keycloakRepo)
	app.Use(middleware.NewJWTMiddleware(keycloakService))

	// Private Routes (Requires Authorization to access these routes)
	privateRouter := app.Group("/")
	KVRouter(privateRouter, kvBaseURL)
	ContactMeRouter(privateRouter)
}
