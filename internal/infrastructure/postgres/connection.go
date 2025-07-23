package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/XxThunderBlastxX/thunder-api/internal/config"
)

type ConnectionManager struct {
	DB *gorm.DB
}

func MustNew(cfg config.DatabaseConfig) *ConnectionManager {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=require&channel_binding=require", cfg.User, cfg.Password, cfg.Host, cfg.Name)

	opt := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(dsn), opt)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	log.Println("🐘 Connected to Database ✅")

	return &ConnectionManager{
		DB: db,
	}
}
