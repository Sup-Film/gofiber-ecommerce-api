package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	APPEnv         string
	APPPort        string
	APPUrl         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPass         string
	DBName         string
	DBSSL          string
	JWTSecret      string
	JWTExpiresIn   string
	AdminEmail     string
	AdminPassword  string
	AdminFirstName string
	AdminLastName  string
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
		DBPass:         getEnv("DB_PASS", ""),
		DBName:         getEnv("DB_NAME", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		AdminEmail:     getEnv("ADMIN_EMAIL", ""),
		AdminPassword:  getEnv("ADMIN_PASSWORD", ""),
		AdminFirstName: getEnv("ADMIN_FIRST_NAME", ""),
		AdminLastName:  getEnv("ADMIN_LAST_NAME", ""),
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
		if config.AdminEmail == "" {
			return fmt.Errorf("ADMIN_EMAIL must be set in production environment")
		}
		if config.AdminPassword == "" {
			return fmt.Errorf("ADMIN_PASSWORD must be set in production environment")
		}
		if config.AdminFirstName == "" {
			return fmt.Errorf("ADMIN_FIRST_NAME must be set in production environment")
		}
		if config.AdminLastName == "" {
			return fmt.Errorf("ADMIN_LAST_NAME must be set in production environment")
		}
	}

	if config.AdminEmail != "" && !isValidEmail(config.AdminEmail) {
		return errors.New("ADMIN_EMAIL is not a valid email address")
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

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	return len(email) > 0 &&
		len(email) <= 254 &&
		strings.Contains(email, "@") &&
		strings.Contains(email, ".") &&
		!strings.HasPrefix(email, "@") &&
		!strings.HasSuffix(email, "@")
}
