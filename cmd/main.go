package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	initDB()
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/get", getTask)
	e.POST("/post", createTask)
	e.PUT("/put/:id", updateTask)
	e.DELETE("/delete/:id", deleteTask)

	e.Start("localhost:8080")
}
