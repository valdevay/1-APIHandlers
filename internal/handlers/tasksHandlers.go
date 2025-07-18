package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {

	tasks, err := h.service.(interface {
		GetAllTasks() ([]taskservice.Task, error)
	}).GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not fetch tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var requestBody taskservice.RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad JSON"})
	}

	task := taskservice.Task{Task: requestBody.Task, IsDone: requestBody.IsDone}
	createdTask, err := h.service.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad task ID"})
	}

	var requestBody taskservice.RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad JSON"})
	}

	task := taskservice.Task{
		Task:   requestBody.Task,
		IsDone: requestBody.IsDone,
	}

	updatedTask, err := h.service.UpdateTask(id, task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad task ID"})
	}

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
