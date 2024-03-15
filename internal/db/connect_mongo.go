package db

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/XxThunderBlastxX/thunder-api/internal/gen/databaseconfig"
)

// TODO: Remove this function
func ConnectMongo(dbConfig *databaseconfig.DatabaseConfig) (*mongo.Database, error) {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dbConfig.Host))
	if err != nil {
		return nil, err
	}
	log.Info("🌐Connected to MongoDB!")

	db := client.Database(dbConfig.DbName)
	return db, nil
}
