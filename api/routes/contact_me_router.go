package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	appConfig "github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func ContactMeRouter(router fiber.Router, config *appConfig.AppConfig) {
	contactMeService := service.NewContactMeService(config.AppConfig.Smtp)
	ctr := controller.ContactMeController{ContactMeService: contactMeService}
	contactMeRoute := router.Group("/contact_me")

	// Middleware
	contactMeRoute.Use(cors.New(cors.Config{
		AllowOrigins: "https://koustav.dev,https://www.koustav.dev",
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
	}))
	contactMeRoute.Use(middleware.VerifyCaptchaToken(config.AppConfig.Cloudflare))

	// Routes
	contactMeRoute.Post("/", ctr.ContactMe())
}
