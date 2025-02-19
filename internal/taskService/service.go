package taskService

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTaskByUserId() ([]Task, error) {
    return s.repo.GetTaskForUser()
}

func (s *TaskService) CreateTask(task Task, userID uint) (Task, error) {
	return s.repo.CreateTask(task, userID)
}

func (s *TaskService) UpdateTask(task Task, id uint) (Task, error) {
	return s.repo.UpdateTask(id, task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
