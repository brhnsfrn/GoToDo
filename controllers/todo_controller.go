package controllers

import (
	"ToDo/models"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var toDos = []models.ToDo{
	{Id: 1, Content: "First ToDo", IsChecked: false, IsDeleted: false},
	{Id: 2, Content: "Second ToDo", IsChecked: false, IsDeleted: false},
}

func GetAll(c *fiber.Ctx) error {
	if len(toDos) <= 0 {
		return c.Status(http.StatusNoContent).JSON(("no data!"))
	}
	return c.Status(http.StatusOK).JSON(toDos)
}

func Post(c *fiber.Ctx) error {
	var newToDo models.ToDo

	if err := c.BodyParser(&newToDo); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}
	newToDo.Id = len(toDos) + 1
	newToDo.IsChecked = false
	newToDo.IsDeleted = false
	toDos = append(toDos, newToDo)
	return c.Status(http.StatusCreated).JSON("Created")
}

func Put(c *fiber.Ctx) error {
	var newToDo models.ToDo

	if err := c.BodyParser(&newToDo); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Bad Request")
	}

	for i, element := range toDos {
		if element.Id == newToDo.Id {
			element.Content = newToDo.Content
			element.IsChecked = newToDo.IsChecked
			toDos = append(toDos[:i], toDos[i+1:]...)
			toDos = append(toDos, element)
		}
	}

	return c.Status(http.StatusOK).JSON("Updated")
}

func Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusNotAcceptable).JSON(err)
	}

	for i, element := range toDos {
		if element.Id == id {
			element.IsDeleted = true
			toDos = append(toDos[:i], toDos[i+1:]...)
			toDos = append(toDos, element)
		}
	}
	return c.Status(http.StatusOK).JSON("Deleted")
}
