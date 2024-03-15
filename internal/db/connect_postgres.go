package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
)

// TODO: Need to fix the connection string
func ConnectPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=postgres user=thunder password=abc@12345 dbname=thunder port=5432 sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Project{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
