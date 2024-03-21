package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/XxThunderBlastxX/thunder-api/internal/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/databaseconfig"
)

func ConnectPostgres(dbConfig *databaseconfig.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Project{}, &domain.TechStack{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
