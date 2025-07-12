package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

}

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Task struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func getTask(c echo.Context) error {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get task"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func createTask(c echo.Context) error {
	var requestBody RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not usefull JSON"})
	}

	task := Task{Task: requestBody.Task, IsDone: requestBody.IsDone}
	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func updateTask(c echo.Context) error {
	id := c.Param("id")

	var requestBody RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Not usefull JSON"})
	}

	var task Task
	if err := db.First(&task, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	task.Task = requestBody.Task
	task.IsDone = requestBody.IsDone
	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}

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
