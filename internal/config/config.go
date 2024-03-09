package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/XxThunderBlastxX/thunder-api/internal/db"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/appconfig"
)

type AppConfig struct {
	AppConfig *appconfig.AppConfig

	Db *mongo.Database

	Timer time.Time
}

func NewAppConfig() *AppConfig {
	config, err := appconfig.LoadFromPath(context.TODO(), "internal/config/config.pkl")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.ConnectMongo(config.Db)
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		AppConfig: config,
		Db:        database,
		Timer:     time.Now(),
	}
}
