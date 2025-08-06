package router

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/server"
)

type Router struct {
	app *server.App
}

func New(app *server.App) *Router {
	return &Router{
		app: app,
	}
}

func (r *Router) RegisterRoutes() {
	r.kvRouter()
	// r.appRouter()
	// r.projectRouter()
	// r.redirectRouter()
}
