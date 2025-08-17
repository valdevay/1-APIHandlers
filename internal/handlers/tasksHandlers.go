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

	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     uint(tsk.ID),
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
			UserId: tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {

	taskRequest := request.Body

	taskToCreate := taskservice.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     uint(createdTask.ID),
		Task:   createdTask.Task,
		IsDone: createdTask.IsDone,
		UserId: createdTask.UserID,
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {

	taskRequest := request.Body

	taskToUpdate := taskservice.Task{
		Task:   taskRequest.Task,
		IsDone: taskRequest.IsDone,
		UserID: taskRequest.UserId,
	}
	updateTask, err := h.service.UpdateTask(int(request.Id), taskToUpdate)

	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     uint(updateTask.ID),
		Task:   updateTask.Task,
		IsDone: updateTask.IsDone,
		UserId: updateTask.UserID,
	}

	return response, nil
}

func (h *TaskHandler) GetUsersUserIdTasks(_ context.Context, request tasks.GetUsersUserIdTasksRequestObject) (tasks.GetUsersUserIdTasksResponseObject, error) {

	userTasks, err := h.service.GetTasksByUserID(request.UserId)
	if err != nil {
		return tasks.GetUsersUserIdTasks404Response{}, nil
	}

	response := tasks.GetUsersUserIdTasks200JSONResponse{}

	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     uint(tsk.ID),
			Task:   tsk.Task,
			IsDone: tsk.IsDone,
			UserId: tsk.UserID,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {

	err := h.service.DeleteTask(int(request.Id))
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksId204Response{}

	return response, nil
}
