package routes

import (
	"ToDo/controllers"

	"github.com/gofiber/fiber/v2"
)

func ToDoRoute(app *fiber.App) {
	app.Get("/", controllers.GetAll)
	app.Post("/", controllers.Post)
	app.Put("/", controllers.Put)
	app.Delete("/:id", controllers.Delete)
}
