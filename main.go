package main

import (
	"ToDo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.ToDoRoute(app)
	app.Listen("0.0.0.0:1938")
}
