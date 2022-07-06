package db

import (
	"ToDo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=todo port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.ToDo{})
	return db
}
