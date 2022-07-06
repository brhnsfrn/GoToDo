package models

import (
	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	Content     string
	IsCompleted bool
}
