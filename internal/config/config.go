package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV                 string
	PORT                string
	MONGO_URI           string
	DB_NAME             string
	PROD_URL            string
	DOCS_URL            string
	CORS_ALLOWED_ORIGIN string
	REDIS_ADDR          string
	POSTGRES_URI        string
	JWT_SECRET          string
}

func (c *Config) IsProduction() bool {
	return c.ENV == "production"
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	var config Config

	if getEnv("ENV", "development") == "development" {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	config.ENV = getEnv("ENV", "development")
	config.PORT = getEnv("PORT", "8080")
	config.DB_NAME = getEnv("DB_NAME", "planify")
	config.MONGO_URI = getEnv("MONGO_URI", "mongodb://root:password@localhost:27017")
	config.PROD_URL = getEnv("PROD_URL", "")
	config.DOCS_URL = getEnv("DOCS_URL", "")
	config.CORS_ALLOWED_ORIGIN = getEnv("CORS_ALLOWED_ORIGIN", "")
	config.REDIS_ADDR = getEnv("REDIS_ADDR", "localhost:6379")
	config.POSTGRES_URI = getEnv("POSTGRES_URI", "postgres://root:password@localhost:5433")
	return &config, nil
}
