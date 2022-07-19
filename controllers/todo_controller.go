package controllers

import (
	"ToDo/models"
	"ToDo/repository"
	"ToDo/responses"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var toDoRepository repository.ToDoRepository = repository.New()

func GetAll(c *fiber.Ctx) error {
	results, err := toDoRepository.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	return c.Status(http.StatusOK).JSON(responses.ToDoResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"ToDos": results}})
}

func Post(c *fiber.Ctx) error {
	var newToDo models.ToDo

	if err := c.BodyParser(&newToDo); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}

	results, err := toDoRepository.Post(newToDo)
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(http.StatusCreated).JSON(responses.ToDoResponse{Status: http.StatusCreated, Message: "Success", Data: &fiber.Map{"ToDo": results}})
}

func Put(c *fiber.Ctx) error {
	var newToDo models.ToDo

	if err := c.BodyParser(&newToDo); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}

	results, err := toDoRepository.Put(newToDo)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(http.StatusOK).JSON(responses.ToDoResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"ToDo": results}})
}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusNotAcceptable).JSON("Bad Request")
	}

	err = toDoRepository.Delete(uint(id))
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(http.StatusOK).JSON(responses.ToDoResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"ToDo": id}})
}
