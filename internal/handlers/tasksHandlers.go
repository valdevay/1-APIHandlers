package handlers

import (
	"context"

	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
	"github.com/valdevay/1-APIHandlers/internal/web/tasks"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     uint(tsk.ID),
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskservice.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     uint(createdTask.ID),
		Task:   createdTask.Task,
		IsDone: createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToUpdate := taskservice.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
	}
	updateTask, err := h.service.UpdateTask(int(request.Id), taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     uint(updateTask.ID),
		Task:   updateTask.Task,
		IsDone: updateTask.IsDone,
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	err := h.service.DeleteTask(int(request.Id))
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.DeleteTasksId204Response{}
	// Просто возвращаем респонс!
	return response, nil
}
