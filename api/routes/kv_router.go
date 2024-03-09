package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/internal/global"
	"github.com/XxThunderBlastxX/thunder-api/internal/repository"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func KVRouter(router fiber.Router, kvBaseURL string) {
	kvRepo := repository.NewKVRepository(kvBaseURL, global.Config.Cloudflare.Token)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.KVController{KVService: kvService}

	kvRoute := router.Group("/kv")

	kvRoute.Get("/:key", ctr.GetValue())
	kvRoute.Get("/", ctr.ListKeys())
	kvRoute.Post("/", ctr.SetKeyValue())
	kvRoute.Delete("/:key", ctr.DeleteKeyValue())
}
