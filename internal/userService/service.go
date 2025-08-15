package userService

import (
	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
)

type UserService interface {
	GetAllUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
	GetUserByID(id int) (User, error)
	GetTasksForUser(userID uint) ([]taskservice.Task, error)
}

type UsersService struct {
	UserRepository
	taskService taskservice.TaskService
}

func NewUserService(repo UserRepository, taskService taskservice.TaskService) UserService {
	return &UsersService{UserRepository: repo, taskService: taskService}
}

func (s *UsersService) GetAllUsers() ([]User, error) {
	return s.UserRepository.GetAllUsers()
}

func (s *UsersService) CreateUser(user User) (User, error) {
	return s.UserRepository.CreateUser(user)
}

func (s *UsersService) UpdateUser(id int, user User) (User, error) {
	return s.UserRepository.UpdateUser(id, user)
}

func (s *UsersService) DeleteUser(id int) error {
	return s.UserRepository.DeleteUser(id)
}

func (s *UsersService) GetUserByID(id int) (User, error) {
	return s.UserRepository.GetUserByID(id)
}

func (s *UsersService) GetTasksForUser(userID uint) ([]taskservice.Task, error) {
	return s.taskService.GetTasksByUserID(userID)
}
