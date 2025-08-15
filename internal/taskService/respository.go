package taskservice

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	UpdateTask(id int, task Task) (Task, error)
	DeleteTask(id int) error
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID uint) ([]Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTask(id int, task Task) (Task, error) {

	var currentTask Task
	if err := r.db.First(&currentTask, id).Error; err != nil {
		return Task{}, err
	}

	currentTask.Task = task.Task
	currentTask.IsDone = task.IsDone
	if err := r.db.Save(&currentTask).Error; err != nil {
		return Task{}, err
	}
	return currentTask, nil
}

func (r *taskRepository) DeleteTask(id int) error {
	return r.db.Delete(&Task{}, id).Error
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
