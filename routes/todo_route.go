package routes

import (
	"ToDo/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/", controllers.GetAll)
	api.Post("/", controllers.Post)
	api.Put("/", controllers.Put)
	api.Delete("/:id", controllers.Delete)
}
