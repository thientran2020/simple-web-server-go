package main

import (
	"log"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Starting server...")

	_, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic("Could not connect to database.db!")
	}
	log.Println("Database connected...!!!")

	// rows, _ := database.Query("SELECT * FROM accounts")
	// var accountId string
	// var userId string
	// var accountType string
	// var balance string
	// var status string

	// for rows.Next() {
	// 	rows.Scan(&accountId, &userId, &accountType, &balance, &status)
	// 	log.Println(accountId + " | " + userId + " | " + accountType + " | " + balance + " | " + status)
	// }

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World...!!!")
	})

	app.Listen(":9090")
}
