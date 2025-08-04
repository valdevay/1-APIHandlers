package userService

import (
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
	GetUserByID(id int) (User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetAllUsers() ([]User, error) {
	users := make([]User, 0)
	err := r.db.Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r *UserRepo) CreateUser(user User) (User, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepo) UpdateUser(id int, user User) (User, error) {
	now := time.Now()
	user.UpdatedAt = now
	err := r.db.Model(&User{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&user).Error
	if err != nil {
		return User{}, err
	}
	return r.GetUserByID(id)
}

func (r *UserRepo) DeleteUser(id int) error {
	now := time.Now()
	return r.db.Model(&User{}).Where("id = ?", id).Update("deleted_at", now).Error
}

func (r *UserRepo) GetUserByID(id int) (User, error) {
	var user User
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return user, err
}
