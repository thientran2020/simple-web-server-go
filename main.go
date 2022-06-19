package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thientran2020/simple-web-server-go/database"
)

func main() {
	database.Connect()

	log.Println("Starting server...")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World...!!!")
	})

	app.Listen(":9090")
}
