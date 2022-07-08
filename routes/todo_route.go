package routes

import (
	"ToDo/controllers"
	"ToDo/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.GetAll)
	api := app.Group("/ToDo", middleware.Authorize())
	api.Post("/", controllers.Post)
	api.Put("/", controllers.Put)
	api.Delete("/:id", controllers.Delete)
}
