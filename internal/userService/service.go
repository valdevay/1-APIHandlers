package userService

type UserService interface {
	GetAllUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
	GetUserByID(id int) (User, error)
}

type UsersService struct {
	UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &UsersService{UserRepository: repo}
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
