package router

import "github.com/XxThunderBlastxX/thunder-api/internal/transport/http/handler"

func (r *Router) kvRouter() {
	grp := r.app.Group("/kv")

	h := handler.NewKVHandler()

	{
		grp.Get("", h.ListKeys())
		grp.Get("/:key", h.GetKey())
		grp.Post("", h.SetKey())
		grp.Delete("", h.DeleteKey())
	}
}
