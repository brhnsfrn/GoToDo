package repository

import (
	"ToDo/db"
	"ToDo/models"
	"fmt"
)

type ToDoRepository struct {
}

func New() ToDoRepository {
	return ToDoRepository{}
}

func (r ToDoRepository) GetAll() ([]models.ToDo, error) {
	var toDos []models.ToDo
	if result := db.DB.Order("created_at desc").Find(&toDos); result.Error != nil {
		fmt.Println(result.Error)
	}
	return toDos, nil
}

func (r ToDoRepository) Post(todo models.ToDo) (models.ToDo, error) {
	if result := db.DB.Create(&todo); result.Error != nil {
		fmt.Println(result.Error)
	}
	return todo, nil
}

func (r ToDoRepository) Put(todo models.ToDo) (models.ToDo, error) {
	if result := db.DB.Model(&todo).Updates(map[string]interface{}{"is_completed": todo.IsCompleted}); result.Error != nil {
		fmt.Println(result.Error)
	}
	return todo, nil
}

func (r ToDoRepository) Delete(id uint) error {
	if result := db.DB.Where("id = ?", id).Delete(&models.ToDo{}); result.Error != nil {
		fmt.Println(result.Error)
	}
	return nil
}
