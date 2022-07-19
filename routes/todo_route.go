package routes

import (
	"ToDo/controllers"
	"ToDo/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Setup(app *fiber.App) {
	app.Get("/", logger.New(), controllers.GetAll)
	api := app.Group("/ToDo", logger.New(), middleware.Authorize())
	api.Post("/", controllers.Post)
	api.Put("/", controllers.Put)
	api.Delete("/:id", controllers.Delete)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "GoToDo Metrics Page"}))
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
