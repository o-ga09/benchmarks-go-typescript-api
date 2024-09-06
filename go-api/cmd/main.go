package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a route on the root path
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the server on port 8080
	app.Listen(":8080")

	// Gracefully shutdown the server
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
