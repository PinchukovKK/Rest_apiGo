package userService

import (	
	"gorm.io/gorm"
	"main.go/internal/taskService"
)

type User struct {
	gorm.Model
	Task []taskService.Task
	Email   string `json:"email"`
	Password string `json:"password"`
}