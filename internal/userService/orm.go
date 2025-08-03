package userservice

type UserORM struct {
	// методы для работы с БД
}

// repository.go
type UserRepository interface {
	GetAll() ([]User, error)
	Create(user User) (User, error)
	Update(id uint, user User) (User, error)
	Delete(id uint) error
}

// service.go
type UserService struct {
	repo UserRepository
}
