package taskservice

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
}

type Task struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id" gorm:"not null"`
}
