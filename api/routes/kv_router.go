package routes

import (
	"fmt"
	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/repository"
	"github.com/XxThunderBlast/thunder-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func KVRouter(router fiber.Router) {
	baseURl := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%v/storage/kv/namespaces/%v", global.Env.CFAccountId, global.Env.KvNamespaceId)

	kvRepo := repository.NewKVRepository(baseURl, global.Env.CFToken)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.KVController{KVService: kvService}

	kvRoute := router.Group("/kv")

	kvRoute.Get("/:key", ctr.GetValue())
	kvRoute.Get("/", ctr.ListKeys())
	kvRoute.Post("/", ctr.SetKeyValue())
	kvRoute.Delete("/:key", ctr.DeleteKeyValue())
}
