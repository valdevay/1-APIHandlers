package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valdevay/1-APIHandlers/internal/db"
	"github.com/valdevay/1-APIHandlers/internal/handlers"
	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
	userservice "github.com/valdevay/1-APIHandlers/internal/userService"
	"github.com/valdevay/1-APIHandlers/internal/web/tasks"
	"github.com/valdevay/1-APIHandlers/internal/web/users"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Tasks
	tasksRepo := taskservice.NewTaskRepository(database)
	tasksService := taskservice.NewTaskService(tasksRepo)
	tasksHandler := handlers.NewTaskHandler(tasksService)

	// Users
	userRepo := userservice.NewUserRepository(database)
	userService := userservice.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Tasks handlers
	tasksStrictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	tasks.RegisterHandlers(e, tasksStrictHandler)

	// Users handlers
	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
