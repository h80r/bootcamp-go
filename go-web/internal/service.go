package internal

type Service interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(name, surname, email string, age int, height float64) (User, error)
	UserExists(id int) (bool, error)
	Modify(id int, name, surname, email string, age int, height float64) (User, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return *users, nil
}

func (s *service) GetById(id int) (User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *service) Create(name, surname, email string, age int, height float64) (User, error) {
	user, err := s.repository.Create(name, surname, email, age, height)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *service) Modify(id int, name, surname, email string, age int, height float64) (User, error) {
	user, err := s.repository.Modify(id, name, surname, email, age, height)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *service) UserExists(id int) (bool, error) {
	_, err := s.repository.GetById(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
