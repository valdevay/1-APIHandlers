package taskservice

type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetTask(id int) (Task, error)
	UpdateTask(id int, task Task) (Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{TaskRepository: repo}
}

func (s *taskService) CreateTask(task Task) (Task, error) {
	return s.TaskRepository.CreateTask(task)
}

func (s *taskService) GetTask(id int) (Task, error) {
	return s.TaskRepository.GetTask(id)
}

func (s *taskService) UpdateTask(id int, task Task) (Task, error) {
	return s.TaskRepository.UpdateTask(id, task)
}

func (s *taskService) DeleteTask(id int) error {
	return s.TaskRepository.DeleteTask(id)
}
