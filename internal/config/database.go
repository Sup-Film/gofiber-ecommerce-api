package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/persistence/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabase
func SetupDatabase(config *Config) *gorm.DB {
	// Database connection string ใช้ข้อมูลจาก config
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBPass,
		config.DBName,
		config.DBPort,
		config.DBSSL,
	)

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL โดยใช้ GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if shouldRunMigration() {
		runMigration(db)

		if err := SeedAdminUser(db, config); err != nil {
			log.Printf("Error seeding admin user: %v\n", err)
		}
	} else {
		autoMigrate := os.Getenv("AUTO_MIGRATE")
		appEnv := os.Getenv("APP_ENV")

		if autoMigrate == "false" {
			log.Printf("Auto migration is disabled. Set AUTO_MIGRATE=true to enable it.")
		} else if appEnv == "production" && autoMigrate != "true" {
			log.Printf("Skipping migration in production environment. Set AUTO_MIGRATE=true to enable it.")
		} else {
			log.Printf("Skipping migration Set AUTO_MIGRATE=true to enable it.")
		}

		// ถ้าไม่ต้องการ migrate ให้ทำการ seed ข้อมูลเริ่มต้น
		if err := SeedAdminUser(db, config); err != nil {
			log.Printf("Error seeding initial data: %v\n", err)
		}
	}

	return db
}

// สร้างฟังก์ชัน ตรวจสอบว่าควร migrate หรือไม่
func shouldRunMigration() bool {

	if os.Getenv("AUTO_MIGRATE") == "false" {
		return false
	}

	if os.Getenv("AUTO_MIGRATE") == "true" {
		return true
	}

	if os.Getenv("APP_ENV") == "development" {
		return true
	}
	return false
}

// ฟังก์ชันสำหรับ migrate
func runMigration(db *gorm.DB) {
	log.Println("Running database migration...")

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
}

func RunMigrationManual(config *Config) error {
	db := SetupDatabase(config)

	log.Println("Running manual database migration...")

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
