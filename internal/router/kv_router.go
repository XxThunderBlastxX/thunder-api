package router

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/handler"
	"github.com/XxThunderBlastxX/thunder-api/internal/service"
)

func (r *Router) kvRouter() {
	g := r.app.Group("/kv")

	s := service.NewKVService(&r.app.Config.Cloudflare)
	h := handler.NewKVHandler(s)

	{
		g.Post("/", h.SetKey())
		g.Get("/:key", h.GetValue())
		g.Get("/", h.ListKeys())
		g.Delete("/:key", h.DeleteKey())
	}
}
