package main

import (
	"ToDo/configs"
	"ToDo/db"
	"ToDo/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func main() {
	configs.SetEnv()
	db.Init()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_DOMAINS"),
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Csrf-Token",
		AllowCredentials: true,
	}))
	app.Use(csrf.New())
	routes.Setup(app)
	log.Fatal(app.Listen(":1938"))
}
