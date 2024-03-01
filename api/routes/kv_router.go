package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/repository"
	"github.com/XxThunderBlast/thunder-api/internal/service"
)

func KVRouter(router fiber.Router) {
	kvRepo := repository.NewKVRepository(kvBaseUrl, global.Env.CFToken)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.KVController{KVService: kvService}

	kvRoute := router.Group("/kv")

	kvRoute.Get("/:key", ctr.GetValue())
	kvRoute.Get("/", ctr.ListKeys())
	kvRoute.Post("/", ctr.SetKeyValue())
	kvRoute.Delete("/:key", ctr.DeleteKeyValue())
}
