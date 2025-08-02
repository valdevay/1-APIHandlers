package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valdevay/1-APIHandlers/internal/db"
	"github.com/valdevay/1-APIHandlers/internal/handlers"
	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
	"github.com/valdevay/1-APIHandlers/internal/web/tasks"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	repo := taskservice.NewTaskRepository(database)
	service := taskservice.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка

	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
