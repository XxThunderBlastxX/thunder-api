package router

import "github.com/XxThunderBlastxX/thunder-api/internal/handler"

func (r *Router) appRouter() {
	h := handler.NewAppHandler()

	r.app.Get("/", h.HandleRequest())
}
