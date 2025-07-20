package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valdevay/1-APIHandlers/internal/db"
	"github.com/valdevay/1-APIHandlers/internal/handlers"
	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := taskservice.NewTaskRepository(database)
	taskService := taskservice.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandler.GetTasks)
	e.POST("/tasks", taskHandler.CreateTask)
	e.PATCH("/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask)

	e.Start("localhost:8080")
}
