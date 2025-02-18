package usersservise

import "gorm.io/gorm"

type UserRepository interface {
	GetUsers() (User, error)
	PostUser(user User) (User, error)
	PatchUserById(id int, updateUser User) (User, error)
	DeleteUserById(id int) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUsers() (User, error) {
	var user User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *Repository) PostUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *Repository) PatchUserById(id int, updateUser User) (User, error) {
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

func (r *Repository) DeleteUserById(id int) error {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}

	err := r.db.Delete(&user).Error
	return err
}