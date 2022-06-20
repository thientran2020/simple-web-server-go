package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/thientran2020/simple-web-server-go/database"
	"github.com/thientran2020/simple-web-server-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	user := models.User{
		Username: data["username"],
		Password: string(password),
		Pin:      data["pin"],
	}

	statement, _ := database.DB.Prepare(`
		INSERT INTO users (username, password, pin) VALUES (?, ?, ?)
	`)
	log.Printf(user.Username + " | " + user.Password + " | " + user.Pin)
	statement.Exec(user.Username, user.Password, user.Pin)

	return c.JSON(user)
}
