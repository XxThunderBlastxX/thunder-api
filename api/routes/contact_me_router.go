package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/XxThunderBlast/thunder-api/api/middleware"
	"github.com/XxThunderBlast/thunder-api/internal/service"
)

func ContactMeRouter(router fiber.Router) {
	contactMeService := service.NewContactMeService()
	ctr := controller.ContactMeController{
		ContactMeService: contactMeService,
	}
	contactMeRoute := router.Group("/contact_me")

	// Middleware
	contactMeRoute.Use(middleware.VerifyCaptchaToken())

	// Routes
	contactMeRoute.Post("/", ctr.ContactMe())
}
