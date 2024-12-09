package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBSource string
	// 他の設定項目をここに追加
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
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
