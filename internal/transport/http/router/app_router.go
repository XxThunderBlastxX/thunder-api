package router

import "github.com/XxThunderBlastxX/thunder-api/internal/transport/http/handler"

func (r *Router) appRouter() {
	h := handler.NewAppHandler()

	r.app.Get("/", h.HandleRequest())
}
