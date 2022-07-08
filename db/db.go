package db

import (
	"ToDo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=12345 dbname=todo port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the PostgreSQL database.")
	}

	DB = db

	db.AutoMigrate(&models.ToDo{})
}
