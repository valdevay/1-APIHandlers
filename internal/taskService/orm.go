package taskservice

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Task struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
