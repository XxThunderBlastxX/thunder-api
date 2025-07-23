package server

import (
	"github.com/XxThunderBlastxX/thunder-api/internal/config"
	"github.com/XxThunderBlastxX/thunder-api/internal/infrastructure/postgres"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	*fiber.App

	Config *config.AppConfig
	PgConn *postgres.ConnectionManager
}

func New() *App {
	// Initializing Fiber app
	app := fiber.New()

	// Initializing app configuration
	cfg := config.NewAppConfig()

	// Initializing PostgreSQL connection
	pgConn := postgres.MustNew(cfg.Database)

	return &App{
		App:    app,
		Config: cfg,
		PgConn: pgConn,
	}
}
