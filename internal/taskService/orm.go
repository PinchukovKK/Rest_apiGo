package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserId uint `json:"user_id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
