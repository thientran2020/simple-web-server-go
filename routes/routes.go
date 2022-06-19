package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thientran2020/simple-web-server-go/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
