package server

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/XxThunderBlastxX/thunder-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	*fiber.App

	Config *config.AppConfig
	PgConn *database.ConnectionManager
}

func New() *App {
	// Initializing Fiber app
	app := fiber.New()

	// Initializing app configuration
	cfg := config.NewAppConfig()

	// Initializing PostgreSQL connection
	pgConn := database.New(cfg.Database)

	return &App{
		App:    app,
		Config: cfg,
		PgConn: pgConn,
	}
}
