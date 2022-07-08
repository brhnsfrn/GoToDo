package main

import (
	"ToDo/db"
	"ToDo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Init()
	app := fiber.New()

	routes.Setup(app)
	app.Listen(":1938")
}
