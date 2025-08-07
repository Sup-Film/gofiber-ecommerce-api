package config

import (
	"log"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/domain/entities"
	"github.com/Sup-Film/fiber-ecommerce-api/pkg/utils"
	"gorm.io/gorm"
)

func SeedAdminUser(db *gorm.DB, config *Config) error {

	// ตรวจสอบว่ามี admin user อยู่แล้วหรือไม่
	var count int64
	db.Model(&models.User{}).Where("role = ?", entities.RoleAdmin).Count(&count)

	if count > 0 {
		log.Println("Admin user already exists, skipping seeding")
		return nil
	}

	// ตรวจสอบว่ามีการตั้งค่า admin credentials หรือไม่
	if config.AdminEmail == "" {
		log.Println("⚠️  ADMIN_EMAIL not set, skipping admin user seeding")
		log.Println("💡 To create admin user, set ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_FIRST_NAME, ADMIN_LAST_NAME in .env")
		return nil
	}

	if config.AdminPassword == "" {
		log.Println("⚠️  ADMIN_PASSWORD not set, skipping admin user seeding")
		return nil
	}

	if config.AdminFirstName == "" {
		log.Println("⚠️  ADMIN_FIRST_NAME not set, skipping admin user seeding")
		return nil
	}

	if config.AdminLastName == "" {
		log.Println("⚠️  ADMIN_LAST_NAME not set, skipping admin user seeding")
		return nil
	}

	// ตรวจสอบความยาวของรหัสผ่าน
	if len(config.AdminPassword) < 8 {
		log.Println("⚠️  ADMIN_PASSWORD must be at least 8 characters long, skipping admin user seeding")
		return nil
	}

	// Hash password
	hashedPassword, err := utils.HashedPassword(config.AdminPassword)
	if err != nil {
		log.Printf("❌ Error hashing admin password: %v", err)
		return err
	}

	// สร้าง admin user
	adminUser := &models.User{
		Email:     config.AdminEmail,
		Password:  hashedPassword,
		FirstName: config.AdminFirstName,
		LastName:  config.AdminLastName,
		Role:      entities.RoleAdmin,
		IsActive:  true,
	}

	if err := db.Create(adminUser).Error; err != nil {
		log.Printf("❌ Error creating admin user: %v", err)
		return err
	}

	log.Printf("✅ Admin user created successfully: %s", config.AdminEmail)
	log.Printf("👤 Name: %s %s", config.AdminFirstName, config.AdminLastName)
	log.Println("⚠️  Please ensure you're using a secure password!")

	return nil
}
