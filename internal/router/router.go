package router

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/middleware"
	"github.com/XxThunderBlastxX/thunder-api/internal/server"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	// Register middleware
	r.app.Use(cors.New())
	r.app.Use(middleware.RequestLogger())
	r.app.Use(middleware.RateLimiter())

	r.kvRouter()
	r.projectRouter()
	r.appRouter()
	r.redirectRouter()
}
