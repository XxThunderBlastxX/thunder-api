package config

import (
	"context"
	"database/sql"
	"github.com/XxThunderBlastxX/thunder-api/internal/utils"
	"log"
	"time"

	"github.com/XxThunderBlastxX/thunder-api/internal/config/gen/appconfig"
	"github.com/XxThunderBlastxX/thunder-api/internal/db"
)

type AppConfig struct {
	AppConfig *appconfig.AppConfig
	Db        *sql.DB
	Timer     time.Time
	Favicon   []byte
}

func NewAppConfig() *AppConfig {
	config, err := appconfig.LoadFromPath(context.TODO(), "internal/config/config.pkl")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.ConnectDb(config.Db)
	if err != nil {
		log.Fatal(err)
	}

	favicon, err := utils.GetFavicon()
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		AppConfig: config,
		Db:        database,
		Timer:     time.Now(),
		Favicon:   favicon,
	}
}
