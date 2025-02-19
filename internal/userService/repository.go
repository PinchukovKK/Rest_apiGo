package userService

import (
	"gorm.io/gorm"
	"main.go/internal/taskService"
)

type UserRepository interface {
	GetUsers() ([]User, error)
	PostUser(user User) (User, error)
	PatchUserById(id int, updateUser User) (User, error)
	DeleteUserById(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error

	for i := range user {
		tasks, err := r.GetTasksForUser(user[i].ID)
		if err != nil {
			return nil, err
		}
		user[i].Task = tasks
	}

	return user, err
}

func (r *userRepository) PostUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) PatchUserById(id int, updateUser User) (User, error) {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}
	if updateUser.Password != "" {
		user.Password = updateUser.Password
	}

	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUserById(id int) error {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	err := r.db.Delete(&user).Error
	return err
}

func (r *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var task []taskService.Task
	if err := r.db.Where("user_id = ?", userID).Find(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}