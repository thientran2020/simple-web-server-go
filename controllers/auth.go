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
	statement.Exec(user.Username, user.Password, user.Pin)
	log.Println("User has been successfully created...!")

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	statement, _ := database.DB.Prepare(`
		SELECT password FROM users WHERE username = ? LIMIT 1
	`)
	rows, err := statement.Query(data["username"])

	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found!!!",
		})
	}

	var password []byte
	for rows.Next() {
		err := rows.Scan(&password)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := bcrypt.CompareHashAndPassword(password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Logged In!!!",
	})
}
