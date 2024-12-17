package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBSource string
}

func Load() *Config {
	// .env.localが存在する場合のみ読み込む
	if _, err := os.Stat(".env.local"); err == nil {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Printf("Warning: Failed to load .env.local: %v", err)
		}
	}

	return &Config{
		Port:     getEnvWithDefault("PORT", "8080"),
		DBSource: getEnvWithDefault("DB_SOURCE", "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable"),
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
