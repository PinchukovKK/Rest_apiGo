package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.PostUser(user)
}

func (s *UserService) PatchUser(id int, updateUser User) (User, error) {
	return s.repo.PatchUserById(id, updateUser)
}	

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUserById(id)
}

func (s *UserService) GetUsers() ([]User, error) {
	return s.repo.GetUsers()
}

