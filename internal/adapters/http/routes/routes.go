package routes

import (
	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/http/middleware"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/gofiber/swagger"
)

// SetupRoutes กำหนดเส้นทาง (routes) สำหรับแอปพลิเคชัน
func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	adminHandler := handlers.NewAdminHandler()

	// Swagger
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	// Api Routes กำหนดกลุ่มเส้นทางสำหรับ API
	api := app.Group("/api")

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protect Routes
	user := api.Group("/user")
	user.Use(middleware.AuthMiddleware())
	user.Get("/profile", authHandler.GetUserProfile)

	// Admin Only Routes
	// กำหนด middleware สำหรับเส้นทางที่ต้องการการยืนยันตัวตนและสิทธิ์
	// ใช้ middleware สำหรับการตรวจสอบสิทธิ์ที่เขียนไว้ในไฟล์ middleware/auth_middleware.go
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	admin.Get("/dashboard", adminHandler.GetDashboard)
}
