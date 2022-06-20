package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thientran2020/simple-web-server-go/database"
	"github.com/thientran2020/simple-web-server-go/routes"
)

func main() {
	database.Connect()

	log.Println("Starting server...")

	app := fiber.New()

	routes.Setup(app)
	routes.Register(app)

	app.Listen(":9090")

	// Some testing
	// More testing
}
