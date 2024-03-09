package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/api/middleware"
	appConfig "github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func ContactMeRouter(router fiber.Router, config *appConfig.AppConfig) {
	contactMeService := service.NewContactMeService(config.AppConfig.Smtp)
	ctr := controller.ContactMeController{
		ContactMeService: contactMeService,
	}
	contactMeRoute := router.Group("/contact_me")

	// Middleware
	contactMeRoute.Use(middleware.VerifyCaptchaToken(config.AppConfig.Cloudflare))

	// Routes
	contactMeRoute.Post("/", ctr.ContactMe())
}
