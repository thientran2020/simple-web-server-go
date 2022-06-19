package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting server...")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World...!!!")
	})

	app.Listen(":9090")
}
