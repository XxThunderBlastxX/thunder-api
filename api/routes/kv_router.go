package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/internal/config/gen/cloudflareconfig"
	"github.com/XxThunderBlastxX/thunder-api/internal/repository"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func KVRouter(router fiber.Router, kvBaseURL string, cfConfig *cloudflareconfig.CloudflareConfig) {
	kvRepo := repository.NewKVRepository(kvBaseURL, cfConfig.Token)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.KVController{KVService: kvService}

	kvRoute := router.Group("/kv")

	kvRoute.Get("/:key", ctr.GetValue())
	kvRoute.Get("/", ctr.ListKeys())
	kvRoute.Post("/", ctr.SetKeyValue())
	kvRoute.Delete("/:key", ctr.DeleteKeyValue())
}
