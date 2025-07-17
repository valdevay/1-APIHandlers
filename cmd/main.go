package main

import (
	"log"

	"example.com/mymodule/internal/db"
	"example.com/mymodule/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/get", taskHandler.GetTasks)
	e.POST("/post", taskHandler.CreateTask)
	e.PUT("/put/:id", taskHandler.UpdateTask)
	e.DELETE("/delete/:id", taskHandler.DeleteTask)

	e.Start("localhost:8080")
}
