package db

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/XxThunderBlast/thunder-api/internal/global"
)

func ConnectMongo() (*mongo.Database, error) {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(global.Config.Db.Host))
	if err != nil {
		return nil, err
	}
	log.Info("🌐Connected to MongoDB!")

	db := client.Database(global.Config.Db.DbName)
	return db, nil
}
