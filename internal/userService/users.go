package userservice

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (m *Migration) CreateUsersTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
