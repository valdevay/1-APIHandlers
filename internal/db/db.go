package db

import (
	"log"

	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	if err := db.AutoMigrate(&taskservice.Task{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
	return db, nil
}
