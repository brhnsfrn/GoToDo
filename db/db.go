package db

import (
	"ToDo/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DBURI")), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the PostgreSQL database.")
	}

	DB = db

	db.AutoMigrate(&models.ToDo{})
}
