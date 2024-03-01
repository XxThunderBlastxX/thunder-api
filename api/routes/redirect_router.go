package routes

import (
	"fmt"
	"github.com/XxThunderBlast/thunder-api/api/controller"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"github.com/XxThunderBlast/thunder-api/internal/repository"
	"github.com/XxThunderBlast/thunder-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func RedirectRouter(router fiber.Router) {
	baseURL := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%v/storage/kv/namespaces/%v", global.Env.CFAccountId, global.Env.KvNamespaceId)

	kvRepo := repository.NewKVRepository(baseURL, global.Env.CFToken)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.RedirectController{KVService: kvService}

	router.Get("/:key", ctr.Redirect())
}
