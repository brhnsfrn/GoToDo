package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"ToDo/repository"
	"ToDo/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB = db.Init()
var h repository.ToDoRepository = repository.New(DB)

func GetAll(c *fiber.Ctx) error {
	results, err := h.GetAll()

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

	results, err := h.Post(newToDo)
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

	results, err := h.Put(newToDo)
	if err != nil {
		fmt.Println(err)
	}

	return c.Status(http.StatusOK).JSON(responses.ToDoResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"ToDo": results}})
}

func Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusNotAcceptable).JSON("Bad Request")
	}

	err = h.Delete(uint(id))
	if err != nil {
		fmt.Println(err)
	}
	return c.Status(http.StatusOK).JSON(responses.ToDoResponse{Status: http.StatusOK, Message: "Success", Data: &fiber.Map{"ToDo": id}})
}
