package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Go Fiber E-commerce API!")
	})

	if err := app.Listen(":3000"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
