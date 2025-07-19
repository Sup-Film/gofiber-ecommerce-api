package config

import (
	"os"
)

type Config struct {
	APPEnv       string
	APPPort      string
	APPUrl       string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPass       string
	DBName       string
	DBSSL        string
	JWTSecret    string
	JWTExpiresIn string
}

func LoadConfig() *Config {
	return &Config{
		APPEnv:       getEnv("APP_ENV", "development"),
		APPPort:      getEnv("APP_PORT", "3000"),
		APPUrl:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPass:       getEnv("DB_PASS", "1234"),
		DBName:       getEnv("DB_NAME", "gofiberecommerce"),
		DBSSL:        getEnv("DB_SSL", "disable"),
		JWTSecret:    getEnv("JWT_SECRET", "your_jwt_secret"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
