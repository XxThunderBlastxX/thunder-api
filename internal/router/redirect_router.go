package router

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/handler"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func (r *Router) redirectRouter() {
	s := service.NewKVService(&r.app.Config.Cloudflare)

	h := handler.NewRedirectHandler(s)

	{
		r.app.Get("/:key", h.Redirect())
	}
}
