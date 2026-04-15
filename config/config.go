package config

import (
	"log"
	"os"
)

type AppConfig struct {
	AppPort   string
	DBUrl     string
	JWTSecret string
}

func LoadConfig() AppConfig {
	return AppConfig{
		AppPort:   getEnv("APP_PORT", "3000"),
		DBUrl:     getEnv("DB_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		if fallback == "" {
			log.Fatalf("Missing required env: %s", key)
		}
		return fallback
	}
	return value
}