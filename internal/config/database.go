package config

import (
	"fmt"
	"log"

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

	// Automatically migrate the schema
	// นำข้อมูลโมเดล User ไปสร้างตารางในฐานข้อมูล
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	// ถ้าเชื่อมต่อสำเร็จ
	log.Println("Database connection established successfully")
	return db
}
