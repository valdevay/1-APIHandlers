package taskservice

type TaskService interface {
	CreateTask(task Task) (Task, error)
	UpdateTask(id int, task Task) (Task, error)
	DeleteTask(id int) error
	GetAllTasks() ([]Task, error)
	GetTasksByUserID(userID uint) ([]Task, error)
}

func (s *TasksService) PatchTask(id int, updatedTask Task) (Task, error) {
	return s.UpdateTask(id, updatedTask)
}

type TasksService struct {
	TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &TasksService{TaskRepository: repo}
}

func (s *TasksService) GetAllTasks() ([]Task, error) {
	return s.TaskRepository.GetAllTasks()
}

func (s *TasksService) UpdateTask(id int, task Task) (Task, error) {
	return s.TaskRepository.UpdateTask(id, task)
}

func (s *TasksService) DeleteTask(id int) error {
	return s.TaskRepository.DeleteTask(id)
}

func (s *TasksService) GetTasksByUserID(userID uint) ([]Task, error) {
	return s.TaskRepository.GetTasksByUserID(userID)
}
