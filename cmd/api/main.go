package main

import (
	"log"

	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/http/handlers"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/http/routes"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/adapters/persistence/repositories"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/config"
	"github.com/Sup-Film/fiber-ecommerce-api/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	// Setup Database
	db := config.SetupDatabase(cfg)

	// เริ่มต้นตั่งค่า Repositories
	userRepo := repositories.NewUserRepository(db)

	// เริ่มต้นตั่งค่า Services
	authService := services.NewAuthService(userRepo)

	// เริ่มต้นตั่งค่า Handlers
	authHandler := handlers.NewAuthHandler(authService)

	// สร้างแอปพลิเคชัน Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup Routes
	routes.SetupRoutes(app, authHandler)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	log.Printf("Server is running on port: %s\n", cfg.APPPort)
}
