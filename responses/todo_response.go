package responses

import "github.com/gofiber/fiber/v2"

type ToDoResponse struct {
	Status  int
	Message string
	Data    *fiber.Map
}
