package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

func LoadConfig() (*Config, error) {
	// โหลดไฟล์ .env ถ้ามี
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	} else {
		log.Println("Loaded .env file successfully")
	}

	config := &Config{
		// ค่าที่ปลอดภัยสำหรับการพัฒนา
		APPEnv:       getEnv("APP_ENV", "development"),
		APPPort:      getEnv("APP_PORT", "3000"),
		APPUrl:       getEnv("APP_URL", "http://localhost:3000"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBSSL:        getEnv("DB_SSL", "disable"),
		JWTExpiresIn: getEnv("JWT_EXPIRES_IN", "24h"),

		// ค่าที่ไม่ปลอดภัยสำหรับการตั่งค่า Default ต้องตั่งค่าในไฟล์ .env เท่านั้น
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}

	// ตรวจสอบค่าที่จำเป็น
	if err := validateConfig(config); err != nil {
		log.Fatalf("Configuration error: %v\n", err)
		return nil, err
	}

	return config, nil
}

// ฟังก์ชันสำหรับตรวจสอบค่า env
func validateConfig(config *Config) error {
	if config.APPEnv == "production" {
		if config.DBPass == "" {
			return fmt.Errorf("DB_PASS must be set in production environment")
		}
		if config.DBName == "" {
			return fmt.Errorf("DB_NAME must be set in production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWTSecret must be set in production environment")
		}
		if config.DBSSL == "disable" {
			return fmt.Errorf("DB_SSL must be enabled in production environment")
		}
	}

	// ตรวจสอบค่าพื้นฐาน
	if config.DBName == "" {
		return fmt.Errorf("DB_NAME must be set")
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
