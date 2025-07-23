package config

import (
	"log"
	"os"
	"time"

	"github.com/XxThunderBlastxX/thunder-api/internal/utils"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name       string
	Version    string
	Port       string
	Cloudflare CloudflareConfig
	Database   DatabaseConfig
	Timer      time.Time
	Favicon    []byte
}

type CloudflareConfig struct {
	AccountID     string
	KVNamespaceID string
	Token         string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var (
	appName    = "thunder-api"
	appVersion = "0.2.0"
	appPort    = "3000"
)

func NewAppConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using environment variables directly")
	}

	cc := CloudflareConfig{
		AccountID:     os.Getenv("CF_ACCOUNT_ID"),
		KVNamespaceID: os.Getenv("CF_KV_NAMESPACE_ID"),
		Token:         os.Getenv("CF_API_TOKEN"),
	}

	dc := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	if port, ok := os.LookupEnv("PORT"); ok {
		appPort = port
	}

	favicon, err := utils.GetFavicon()
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		Name:       appName,
		Version:    appVersion,
		Port:       appPort,
		Cloudflare: cc,
		Database:   dc,
		Timer:      time.Now(),
		Favicon:    favicon,
	}
}
