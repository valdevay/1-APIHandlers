package userService

import (
	"time"

	taskservice "github.com/valdevay/1-APIHandlers/internal/taskService"
)

type User struct {
	ID        int                `json:"id" gorm:"primaryKey"`
	Email     string             `json:"email" gorm:"uniqueIndex;not null"`
	Password  string             `json:"password" gorm:"not null"`
	CreatedAt time.Time          `json:"-" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time          `json:"-" gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time         `json:"-" gorm:"index"`
	Tasks     []taskservice.Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
