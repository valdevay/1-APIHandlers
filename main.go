package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type RequestBody struct {
	Task string `json:"task"`
}

var task string

func getTask(c echo.Context) error {
	response := fmt.Sprintf("hello, %s", task)
	return c.String(http.StatusOK, response)
}

func createTask(c echo.Context) error {
	var requestBody RequestBody

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный формат JSON"})
	}

	task = requestBody.Task

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"task": requestBody.Task,
	})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/get", getTask)
	e.POST("/post", createTask)

	e.Start("localhost:8080")
}
