package main

import (
	"ToDo/configs"
	"ToDo/db"
	"ToDo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.SetEnv()
	db.Init()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":1938")
}
