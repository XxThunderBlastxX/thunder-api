package db

import (
	"database/sql"
	"fmt"

	"github.com/XxThunderBlastxX/thunder-api/internal/config/gen/databaseconfig"
)

func ConnectDb(dbConfig *databaseconfig.DatabaseConfig) (*sql.DB, error) {
	dsn := dbConfig.Dsn

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("🎉Database connected !!!")

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	fmt.Println("✅ Database ping successfully !!!")

	return db, nil
}
