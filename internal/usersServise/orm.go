package usersservise

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string `json:"email"`
	Password uint `json:"password"`
}