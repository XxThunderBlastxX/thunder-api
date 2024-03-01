package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/repository"
	"github.com/XxThunderBlast/thunder-api/internal/service"
)

func RedirectRouter(router fiber.Router) {
	kvRepo := repository.NewKVRepository(kvBaseUrl, global.Env.CFToken)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.RedirectController{KVService: kvService}

	router.Get("/:key", ctr.Redirect())
}
