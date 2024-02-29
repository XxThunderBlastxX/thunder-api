package routes

import (
	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/repository"
	"github.com/XxThunderBlast/thunder-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RedirectRouter(router fiber.Router) {
	kvRepo := repository.NewKVRepository(global.BaseKVPath, global.Env.CFToken)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.RedirectController{KVService: kvService}

	router.Get("/:key", ctr.Redirect())
}
