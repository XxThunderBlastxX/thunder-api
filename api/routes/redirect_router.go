package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XxThunderBlastxX/thunder-api/api/controller"
	"github.com/XxThunderBlastxX/thunder-api/internal/global"
	"github.com/XxThunderBlastxX/thunder-api/internal/repository"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func RedirectRouter(router fiber.Router, kvBaseURL string) {
	kvRepo := repository.NewKVRepository(kvBaseURL, global.Config.Cloudflare.Token)
	kvService := service.NewKVService(kvRepo)
	ctr := controller.RedirectController{KVService: kvService}

	router.Get("/:key", ctr.Redirect())
}
