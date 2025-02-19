package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task, userID uint) (Task, error)
	GetAllTask() ([]Task, error)
	UpdateTask(id uint, task Task) (Task, error)
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task, userID uint) (Task, error) {
	task.UserId = userID
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTask() ([]Task, error) {
	var tasks []Task
    err := r.db.Find(&tasks).Error
    return tasks, err
}

func (r *taskRepository) UpdateTask(id uint, updatedTask Task) (Task, error) {
	var task Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}

	if updatedTask.Task != "" {
		task.Task = updatedTask.Task
	}
	task.IsDone = updatedTask.IsDone

	err := r.db.Save(&task).Error
	return task, err
}

func (r *taskRepository) DeleteTask(id uint) error {
	var task Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return result.Error
	}

	err := r.db.Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}
