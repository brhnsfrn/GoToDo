package main

import (
	"ToDo/configs"
	"ToDo/db"
	"ToDo/routes"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func main() {
	configs.SetEnv()
	db.Init()
	fmt.Println(500 * time.Millisecond)
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5500, http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Csrf-Token",
		AllowCredentials: true,
	}))
	app.Use(csrf.New())
	routes.Setup(app)
	log.Fatal(app.Listen(":1938"))
}
