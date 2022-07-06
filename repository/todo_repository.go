package repository

import (
	"ToDo/models"
	"fmt"

	"gorm.io/gorm"
)

type ToDoRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) ToDoRepository {
	return ToDoRepository{db}
}

func (r ToDoRepository) GetAll() ([]models.ToDo, error) {
	var toDos []models.ToDo
	if result := r.db.Find(&toDos); result.Error != nil {
		fmt.Println(result.Error)
	}
	return toDos, nil
}

func (r ToDoRepository) Post(todo models.ToDo) (models.ToDo, error) {
	if result := r.db.Create(&todo); result.Error != nil {
		fmt.Println(result.Error)
	}
	return todo, nil
}

func (r ToDoRepository) Put(todo models.ToDo) (models.ToDo, error) {
	if result := r.db.Model(&todo).Updates(map[string]interface{}{"content": todo.Content, "is_completed": todo.IsCompleted}); result.Error != nil {
		fmt.Println(result.Error)
	}
	return todo, nil
}

func (r ToDoRepository) Delete(id uint) error {
	if result := r.db.Where("id = ?", id).Delete(&models.ToDo{}); result.Error != nil {
		fmt.Println(result.Error)
	}
	return nil
}
