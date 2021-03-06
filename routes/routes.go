package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thientran2020/simple-web-server-go/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}

func Register(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
}
