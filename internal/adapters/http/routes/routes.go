package routes

import (
	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// SetupRoutes กำหนดเส้นทาง (routes) สำหรับแอปพลิเคชัน
func SetupRoutes(app *fiber.App, authHandler *handlers.AuthHandler) {
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Api Routes กำหนดกลุ่มเส้นทางสำหรับ API
	api := app.Group("/api")

	// Auth Routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register) // Routes ไปเรียกใช้ Handler สำหรับการลงทะเบียนผู้ใช้
	auth.Post("/login", authHandler.Login)       // Routes ไปเรียกใช้ Handler สำหรับการเข้าสู่ระบบผู้ใช้

	// Protected Routes
	admin := api.Group("/admin")
	// กำหนด middleware สำหรับเส้นทางที่ต้องการการยืนยันตัวตนและสิทธิ์
	// ใช้ middleware สำหรับการตรวจสอบสิทธิ์ที่เขียนไว้ในไฟล์ middleware/auth_middleware.go
	admin.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the admin dashboard",
		})
	})
}
